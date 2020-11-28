package main

import "os"

//file opening

func openFile(fileName string) (*os.File, error) {
	file, err := os.OpenFile("last_update.txt", os.O_RDWR, os.ModeExclusive)

	return file, err
}
