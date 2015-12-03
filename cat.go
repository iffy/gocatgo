package main

import (
	"fmt"
	"os"
	"log"
	"io/ioutil"
)

func main() {
	files := os.Args[1:]

	for _, file := range files {
		data, err := ioutil.ReadFile(file)
		if err != nil {
			log.Printf("Error reading file: %v %v", file, err)
		}
		fmt.Printf("%s", data)
	}
}
