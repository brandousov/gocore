// App Core functions and vars | Compression functions
package core

import (
	"bufio"
	"compress/gzip"
	"io/ioutil"
	"os"
)





////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Сжатие указанного файла
// @help: https://www.dotnetperls.com/compress-go
func GzFile(filePath string) error {
	f1, err1		:=	os.Open(filePath)
	if err1 != nil { return err1 }

	reader			:=	bufio.NewReader(f1)
	content, err2	:=	ioutil.ReadAll(reader)
	if err2 != nil { return err2 }

	f2, err3	:=	os.Create(filePath +".gz")
	if err3 != nil { return err3 }

	w, err4 :=	gzip.NewWriterLevel(f2, gzip.BestCompression)
	if err4 != nil { return err4 }

	_, _	=	w.Write(content)
	_		=	w.Close()
	_ = f2.Close()
	_ = f1.Close()

	return nil
}
