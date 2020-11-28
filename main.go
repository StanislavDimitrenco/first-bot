package main

import (
	_ "github.com/mattn/go-sqlite3"
	"time"
)

func main() {
	file, err := openFile("last_update.txt")
	checkError(err)
	defer file.Close()

	offset64, err := stringToInt(byteToString(file))
	checkError(err)

	offset := int(offset64)

	engine(MethodGetUpdates, offset, file)
	time.Sleep(1000 * time.Millisecond)
}
