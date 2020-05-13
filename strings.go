// Bioio API v7 © 2019 ITCorp (it.ru)
//
// App Core functions and vars
package core

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"html"
	"net"
	"strconv"
	"strings"
)





//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////// |strings|
// Преобразование текста в ВЕРХНИЙ регистр
func Str2Upper(s string) string {
	return strings.ToUpper(s)
}





// Преобразование текста в нижний регистр
func Str2Lower(s string) string {
	return strings.ToLower(s)
}





// [+] Упаковка в JSON | data map[string]interface{}
// @help: https://golang.org/pkg/encoding/json/#Marshal
// @help: https://stackoverflow.com/questions/30105798/go-constructing-struct-json-on-the-fly
func Json(data interface{}) string {
	result, err := json.Marshal(data)
	if err != nil {
		Out("core/strings.JsonAry() error: %#v", err)
		return ""
	}
	return string(result)
}





// [+] Распаковка JSON строки
// @help: https://stackoverflow.com/questions/33436730/unmarshal-json-with-some-known-and-some-unknown-field-names
func DeJson(s string) interface{} {
	var result interface{}
	if err := json.Unmarshal([]byte(s), &result); err != nil {
		Out("core/strings.DeJson() failed: %#v", err)
		return nil
	}
	return result
}





// Заполнение ведущими нолями
func Zerofill(v int, zeros int, total int) string {
	vlen	:=	len(string(v))
	format	:=	fmt.Sprintf("%%0%dd", zeros + vlen)
	str		:=	fmt.Sprintf(format, v)
	return str[len(str) - total:]
}





// Очистка пробелов в начале и конце строки
// @help: https://golang.org/pkg/strings/#TrimSpace
func Trim(v string) string {
	return strings.TrimSpace(v)
}





// Форматирование строки
func Sprintf(format string, v ...interface{}) string {
	return fmt.Sprintf(format, v...)
}





// Вытянуть массив строк в одну строку с разделителем
// @help: https://stackoverflow.com/questions/28799110/how-to-join-a-slice-of-strings-into-a-single-string
func Join(v []string, delimiter string) string {
	return strings.Join(v, delimiter)
}





// Преобразование числа в строку
// @help: https://golang.org/pkg/strconv/
func Int2str(v int64) string {
	return strconv.FormatInt(v, 10)
}





// Преобразование числа в строку
// @help: https://golang.org/pkg/strconv/
func Uint2str(v uint64) string {
	return strconv.FormatUint(v, 10)
}





// Преобразование строки в число
// @help: https://golang.org/pkg/strconv/
func Str2int(v string) int64 {
	res, _	:=	strconv.ParseInt(v, 10, 64)
	return res
}





// Экранирование спецсимволов HTML
// @help: https://golang.org/pkg/html/#example_EscapeString
// var htmlReplacer = strings.NewReplacer(
//	"&", "&amp;",
//	"<", "&lt;",
//	">", "&gt;",
//	"&#34;" is shorter than "&quot;".
//	`"`, "&#34;",
//	"&#39;" is shorter than "&apos;" and apos was not in HTML until HTML5.
//	"'", "&#39;",
//)
func Hs(v string) string {
	return html.EscapeString(v)
}





// Преобразование IP адреса в число
// @help: https://github.com/go-libs/iputils/blob/master/iputils.go
func IP2Long(v string) int64 {
	ip := net.ParseIP(v)
	if ip == nil { return 0 }
	ip = ip.To4()
	return int64(binary.BigEndian.Uint32(ip))
}





// Преобразование числа в IP адрес
func Long2IP(v uint32) string {
	ipByte := make([]byte, 4)
	binary.BigEndian.PutUint32(ipByte, v)
	ip := net.IP(ipByte)
	return ip.String()
}
