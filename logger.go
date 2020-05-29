// App data Types | App Logger{} struct
package core

import (
	"fmt"
	"os"
	"sync"
)





//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////// |logger|
// Logger struct
type Logger struct {
	File *os.File   // открытый лог-файл для записи лога
	Dir  string     // директория лог-файла
	Name string     // имя/название лог-файла
	RPS  int        // кол-во запросов в секунду
	Ary  []string   // кэш для записей
	Mux  sync.Mutex // mutex для стопа конкурентных операций с логами
	Dbg  bool       // Флаг вывода отладочной информации
	Ymd  string     // Текущая метка ГГГГДДММ (обновляется ежесекундно)
	//StopRotate	chan bool																								// Канал для сигнала остановки горутины с ротацией лога
}





//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////// |methods|
// Запись в логфайл
func (lgr *Logger) Write(log string, console bool) {
	if log == "" {
		return
	}
	log = LogDatePrefix() + log + "\n"
	if lgr.File != nil {
		_, _ = lgr.File.Write([]byte(log))
	}

	// Если заказан вывод логов в консоль
	if console {
		fmt.Print(log)
	} else {
		// Вывод в консоль делаем только при дебаге!
		if lgr.Dbg {
			fmt.Print(log)
		}
	}
}





// Запись HTTP-запросов в лог-файл
// Адаптировано под Hi-Load, так как частая запись на диск
// очень сильно тормозит скорость обработки HTTP-запросов
func (lgr *Logger) Add(log string) {
	log = LogDatePrefix() + log + "\n"
	lgr.Mux.Lock()
	lgr.RPS++
	lgr.Ary = append(lgr.Ary, log)
	logLen := len(lgr.Ary)
	lgr.Mux.Unlock()

	if logLen > lgr.RPS {
		lgr.WriteCacheToDisk()
	}
}

// Сброс кэша из памяти на диск
func (lgr *Logger) WriteCacheToDisk() {
	lgr.Mux.Lock()
	logLen := len(lgr.Ary)

	if logLen > 0 {
		logs := Join(lgr.Ary, "")
		lgr.Ary = nil
		lgr.Mux.Unlock()
		if lgr.File != nil {
			_, _ = lgr.File.Write([]byte(logs))
		}
	} else {
		lgr.Mux.Unlock()
	}
}





// Закрытие логера
func (lgr *Logger) Close() {
	// записать на диск кэш
	lgr.WriteCacheToDisk()
	lgr.Write("***", false)
	_ = lgr.File.Close()
}





// Запускает ротацию лога, раз в секунду запуская *Logger.Rotate()
func (lgr *Logger) Rotate() {
	RunEvery(1000, lgr.Check)
}





// Проверка даты и времени в именах лог-фалов (выполняется каждую 1 сек в отдельном потоке)
// а также компрессия старых: https://www.dotnetperls.com/compress-go
func (lgr *Logger) Check() {
	// Проверка смены даты для своевременного переключения со старого логфайла приложения
	// на новый с обновлённой датой, после чего старые файлы архивируются и удаляются
	gzFlag := false
	lgr.Ymd = Date("Ymd")
	logPath := lgr.Dir + "/" + lgr.Ymd + "_" + lgr.Name + ".log"

	// Сброс счётчика запросов в секунду для Hi-Load варианта с кэшированным логом
	// *Logger.Check() из *Logger.Rotate() вызывается раз в 1 сек.
	lgr.RPS = 0

	// Проверка признаков "инициализации приложения" или "смены даты"
	if lgr.File == nil || logPath != lgr.File.Name() {
		lgr.Close()
		logFile, err1 := FileOpenForAppend(logPath)
		if err1 != nil {
			lgr.Write(Sprintf("Logger.Rotate() open <%s> logfile error - %#v", lgr.Name, err1), true)
			return
		}
		lgr.File = logFile

		// Все лог-файлы успешно переключились на новую дату
		// теперь можно запускать архивирование старых логов
		gzFlag = true
	}

	// Gzip
	if gzFlag {
		lgr.Gz()
	}
}





// Компрессия старых лог-файлов
func (lgr *Logger) Gz() {
	// Защита от повторного вызова в случае задержки при архивации больших логов
	lgr.Mux.Lock()
	defer lgr.Mux.Unlock()

	// Актуальная дата логфайла
	logYmd := Date("Ymd")

	// Поиск всех логфайлов в папке с логами
	// В случае ошибки - запись в логфайл и разблокировать мутекс
	logFiles, err := Glob(lgr.Dir + "/*" + lgr.Name + ".log")
	if err != nil {
		lgr.Write(Sprintf("Logger.Gz() <%s.log> error - %#v", lgr.Name, err), false)
		return
	}

	// Если не найдено ни одного логфайла
	// запись в логфайл и разблокировать мутекс
	if logFiles == nil {
		lgr.Write("Logger.Gz("+lgr.Name+") log files not found", false)
		return
	}

	// Перебор найденных файлов и работа над ними
	for _, logPath := range logFiles {
		// Сжимать все логи, кроме активных!
		logName := Basename(logPath)
		if logName != logYmd+"_"+lgr.Name+".log" {
			// Сжатие лог-файла
			err := GzFile(logPath)
			if err != nil {
				lgr.Write(Sprintf("Logger.Gz() log file '%s' was NOT compressed, error - %+v", logPath, err), true)
			} else {
				_ = FileDelete(logPath)
				lgr.Write(Sprintf("Logger.Gz() log file '%s' was compressed, archived and deleted", logPath), true)
			}
		}
	}
}





//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////// |functions|
// Создание нового экземпляра Logger{}
func NewLogger(name, LogDir string) *Logger {
	// Создать директорию для логов
	found := FileExists(LogDir)
	if !found {
		if os.MkdirAll(LogDir, os.ModePerm) != nil {
			fmt.Printf(`Logger.NewLogger() - can not create log directory <%s>!`, LogDir)
			return nil
		}
	}

	// Создать логгер и прописать ему актуальную дату и директорию
	logger := Logger{
		Name: name,
		Dir:  LogDir,
		Ymd:  Date("Ymd"),
	}

	// Инициализация и открытие логфайла
	logger.Check()
	return &logger
}





// Возвращает префикс с датой для каждой строки логфайла
func LogDatePrefix() string {
	return Date("Ymd-H:i:s\t")
}
