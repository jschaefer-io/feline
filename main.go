package main

import (
	"fmt"
)

func main() {
	_, err := prepareFile("./testfile.feline")
	if err != nil {
		panic(err)
	}
}

func prepareFile(path string) (File, error) {
	file := NewFile(path)
	err := file.Prepare()
	if err != nil {
		return File{}, err
	}
	for _, line := range file.lines {
		fmt.Printf("Line Number %d:", line.number)
		for _, token := range line.tokens {
			fmt.Println(token)
		}
	}
	//fmt.Println(file)
	return file, nil
}
