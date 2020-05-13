// Bioio API v7 © 2020 ITCorp (it.ru)
//
// Application encryption functions
// @help: GOST https://github.com/martinlindhe/gogost/tree/master/gost3412
package core

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"errors"
)





//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////// |fn|
// Encode bytes to sha1 hex-encoded string
func Md5(input []byte) string {
	hash := md5.New()
	hash.Write(input)
	return Sprintf("%x", hash.Sum(nil))
}





// Encode bytes to sha1 hex-encoded string
func Sha1(input []byte) string {
	hash := sha1.New()
	hash.Write(input)
	return Sprintf("%x", hash.Sum(nil))
}





// Encode bytes to sha256 hex-encoded string
func Sha256(input []byte) string {
	hash := sha256.New()
	hash.Write(input)
	return Sprintf("%x", hash.Sum(nil))
}





// Encode bytes to sha512 hex-encoded string
func Sha512(input []byte) string {
	hash := sha512.New()
	hash.Write(input)
	return Sprintf("%x", hash.Sum(nil))
}





// Encode bytes to base64 encoded string
func Base64(input []byte) string {
	result := base64.StdEncoding.EncodeToString(input)
	return result
}





// Decode base64 encoded string back to []byte array
func Base64Decode(input string) ([]byte, error) {
	result, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		return nil, errors.New(Sprintf("core/crypt.DeBase64() error - %#v", err))
	}
	return result, nil
}





// Encode bytes to hex encoded string
// @help: https://golang.org/pkg/encoding/hex/#EncodeToString
func Bin2hex(input []byte) string {
	out := hex.EncodeToString(input)
	return out
}





// Encode hex encoded string to []byte
// @help: https://golang.org/pkg/encoding/hex/#DecodeString
func Hex2bin(input string) ([]byte, error) {
	out, err := hex.DecodeString(input)
	if err != nil {
		return nil, errors.New(Sprintf("core/crypt.Hex2bin() error - ", err))
	}
	return out, nil
}





// Encode bytes to URL-encoded string
func UrlEncode(input []byte) string {
	result := base64.URLEncoding.EncodeToString(input)
	return result
}





// Decode URL-encoded string back to []byte array
func UrlDecode(input string) ([]byte, error) {
	result, err := base64.URLEncoding.DecodeString(input)
	if err != nil {
		return nil, errors.New(Sprintf("core/crypt.UrlDecode() error - %#v", err))
	}
	return result, nil
}
