// App Core functions and vars
package core

import (
	"net"
)





//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////// |net|
// Закрытие соединения net.Conn, открытого с помощью net.Dial() и тп
func NetConnClose(c net.Conn) bool {
	err	:=	c.Close()
	if err != nil { return false }
	return true
}





// Получение рабочего IPv4 адреса хоста
// @help: https://stackoverflow.com/questions/23558425/how-do-i-get-the-local-ip-address-in-go
// @help: 193.0.14.129 - k.root-servers.org DNS server in Russia, Moscow (https://www.nic.ru/whois/?searchWord=193.0.14.129 - Country NL)
func GetIpv4() (ip net.IP) {
	udp, _	:=	net.Dial("udp", "193.0.14.129:53")
	defer NetConnClose(udp)
	addr	:=	udp.LocalAddr().(*net.UDPAddr)
	return addr.IP
}
