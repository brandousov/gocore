// Bioio API v7 © 2019 ITCorp (it.ru)
//
// App Core functions and vars | Core functions
// @help: in_array: https://github.com/SimonWaldherr/golang-examples/blob/master/advanced/in_array.go
package core

import (
	"fmt"
	"reflect"
	"runtime"
)





//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////// |helpers|
// Вывод значения в консоль во внутреннем формате Go
// @help: https://golang.org/pkg/fmt/#Printf
func Cl(v ...interface{}) {
	fmt.Printf("%#v\n", v)
}





// Вывод значения в консоль
func Out(format string, v ...interface{}) {
	record	:=	format
	if v != nil { record = fmt.Sprintf(format, v...) }
	fmt.Printf("%+v", record)
}





// Возврат читаемого значения полученной переменной
func Ve(v ...interface{}) string {
	return Sprintf("%#v", v)
}





// Проверка типа переменной
func Typeof(v interface{}) string {
	return Sprintf("%T", v)
}





// Проверка наличия элемента в массиве
func IsSet(arr interface{}, index int) bool {
	ary := reflect.ValueOf(arr)
	return ary.Len() > index
}





// Проверка значения на пустую строку, false, 0 или nil
func Empty(v interface{}) bool {
	switch t := v.(type) {
	case string:	return t == ""
	case int:		return t == 0
	case bool:		return t == false
	case nil:		return true
	//case *interface{}:	return t == nil
	default:		return false
	}
}





// Функция для замены отлова паник и фэйлов
// @help: https://golang.org/doc/effective_go.html#recover
func Recover() {
	if err := recover(); err != nil {
		Out("core/core.Recover() handles failing goroutine task, catched error details - %#v", err)
	}
}





// Goexit terminates the goroutine that calls it. No other goroutine is affected.
// Goexit runs all deferred calls before terminating the goroutine.
// Because Goexit is not a panic, any recover calls in those deferred functions will return nil.
// @help: https://golang.org/pkg/runtime/#Goexit
func GoStop() {
	runtime.Goexit()
}
