package main

import (
	"io"
	"os"
	"strconv"
)

func byteToString(file *os.File) string {
	data := make([]byte, 64)
	var str string
	for {
		n, err := file.Read(data)
		if err == io.EOF {
			break
		}
		str = string(data[:n])
	}

	return str
}

func intToByte(i int64) []byte {
	some := strconv.FormatInt(i, 10)
	return []byte(some)
}

func stringToInt(str string) (int64, error) {
	return strconv.ParseInt(str, 0, 64)
}

func overWrite(file *os.File, sls []byte) error {
	_, err := file.WriteAt(sls, 0)
	return err
}

func offsetUpdate(o int, file *os.File) (int, error) {
	o++
	o64 := int64(o)
	return o, overWrite(file, intToByte(o64))
}
