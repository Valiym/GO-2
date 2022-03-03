package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	err := os.MkdirAll("dir", os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 1000000; i++ {
		err := createFile(fmt.Sprintf("%s/%d.txt", "dir", i))
		if err != nil {
			return
		}
	}
}

func createFile(filepath string) (err error) {
	file, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("can't create file: %w", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)
	return
}
