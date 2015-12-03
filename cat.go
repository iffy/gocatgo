package main

import (
	"os"
	"io"
	"log"
)

func main() {
	files := os.Args[1:]
	var rc = 0

	for _, file := range files {
		fh, err := os.Open(file)
		if err != nil {
			log.Printf("Error reading file: %v %v", file, err)
			rc = 1
			continue
		}
		defer fh.Close()
		_, err = io.Copy(os.Stdout, fh)
		if err != nil {
			log.Printf("Error copying file: %v %v", file, err)
			rc = 1
		}
	}
	os.Exit(rc)
}
