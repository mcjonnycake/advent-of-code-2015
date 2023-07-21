package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

var inputPath string = "input.txt"

func getInput(inputPath string) string {
	content, err := ioutil.ReadFile(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}

func main() {
	fmt.Println("ASDF")

	input := getInput(inputPath)

	fmt.Println(input)
}
