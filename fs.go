// App Core functions and vars
package core

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)





//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////// |io.file|
// Чтение файла
func FileRead(filename string) []byte {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("fileRead("+ filename +")", err)
		return []byte("")
	}
	return dat
}





// Запись файла | os.FileMode = 0644
// @todo: сначала надо писать в filename.tmp файл, после чего если места на диске хватило переименовывать его в filename
func FileSave(filename string, data []byte) bool {
	err := ioutil.WriteFile(filename, data, 0644)
	return err == nil
}





// Запись файла | os.FileMode = 0644
// @example: FileAppend("app.log", []byte("Test test test"))
func FileAppend(filename string, data []byte) bool {
	f, err	:=	os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		Out("core/fs.FileAppend() error - %#v", err)
		return false
	}
	defer FileClose(f)
	_, err = f.Write(data)
	if err != nil { return false }
	return true
}





// Проверка наличия файла
// @help: https://stackoverflow.com/questions/12518876/how-to-check-if-a-file-exists-in-go
func FileExists(filename string) bool {
	_, err	:=	os.Stat(filename)
	return err == nil
}





// Получение информации о файле
func FileInfo(filename string) (os.FileInfo, error) {
	return os.Stat(filename)
}





// Удаление файла
func FileDelete(filename string) error {
	return os.Remove(filename)
}





// Открытие файла для дозаписи
func FileOpenForAppend(filePath string) (*os.File, error) {
	_file, err	:=	os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return nil, errors.New(Sprintf("fn/file.FileOpenForAppend() open file <%s> error - %#v", _file, err))
	}
	return _file, nil
}





// Закрытие ранее открытого через os.OpenFile() файла
func FileClose(file *os.File) bool {
	err := file.Close()
	if err != nil {
		return false
	}
	return true
}





// Получение списка файлов в папке по маске
// @help: https://golang.org/pkg/path/filepath/#Glob
func Glob(pattern string) (matches []string, err error) {
	return filepath.Glob(pattern)
}





// Возвращает имя файла
func Basename(path string) string {
	return filepath.Base(path)
}





// Возвращает рабочую директорию приложения
func DirCreate(dir string) bool {
	return os.MkdirAll(dir, os.ModePerm) == nil
}





// Возвращает рабочую директорию приложения
func GetCwd() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		dir = ""
	}
	return dir
}
