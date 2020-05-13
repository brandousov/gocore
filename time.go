// App Core functions and vars
package core

import (
	"strconv"
	"strings"
	"time"
)





//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////// |fn|
// Получение отформатированной даты и времени для вставки в логи и текст
// @help: https://programming.guide/go/format-parse-string-time-date-example.html
// @help: time.Now().Format("20060102 15:04:05")
// @help: date formats - https://www.php.net/manual/ru/function.date.php
func Date(format string, uts ...int32) string {
	_ts		:=	Ts()
	//loc		:=	time.FixedZone("UTC", 0*60*60)

	if len(uts) > 0 { _ts = uts[0] }
	if format == "" { format = "Y-m-d H:i:s" }

	t		:=	time.Unix(int64(_ts), 0).UTC()
	year	:=	strconv.FormatInt(int64(t.Year()), 10)
	month	:=	Zerofill(int(t.Month()), 1, 2)
	day		:=	Zerofill(t.Day(), 1, 2)
	hour	:=	Zerofill(t.Hour(), 1, 2)
	min		:=	Zerofill(t.Minute(), 1, 2)
	sec		:=	Zerofill(t.Second(), 1, 2)

	format	=	strings.ReplaceAll(format, "Y", year)
	format	=	strings.ReplaceAll(format, "m", month)
	format	=	strings.ReplaceAll(format, "d", day)
	format	=	strings.ReplaceAll(format, "H", hour)
	format	=	strings.ReplaceAll(format, "i", min)
	format	=	strings.ReplaceAll(format, "s", sec)

	return format
}





// Получение Unix timestamp
func Ts() int32 {
	return int32(time.Now().Unix())
}





// Sleep в течение ms милисекунд
func Sleep(ms int64) {
	time.Sleep(time.Duration(ms) * time.Millisecond)
}





// Запуск функции через seconds сек в отдельном потоке
// @help: http://golang.org/pkg/time/#AfterFunc
// @help: https://golang.org/pkg/time/#Timer
// @ex: go runAfter(1000, func() { out(`runAfter()`) })
// defer test.Stop()
func RunAfter(ms int64, f func()) {
	var timer *time.Timer
	timer	=	time.AfterFunc(time.Duration(ms) * time.Millisecond, func() {
		timer.Stop()
		f()
	})
}





// Запуск функции каждые seconds сек в отдельном потоке!: go runEvery()
//Tick неуправляем - его невозможно остановить, управление через внешний done <- true
// @help: https://golang.org/pkg/time/#NewTicker
func RunEvery(ms int64, function func()) chan bool {
	ticker	:=	time.NewTicker(time.Duration(ms) * time.Millisecond)
	done	:=	make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				ticker.Stop()
				return
			case <-ticker.C:
				function()
			}
		}
	}()
	return done
}
