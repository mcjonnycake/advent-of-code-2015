package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

var test_file string = "input.txt"

var text string = "Hello, World!"
var test_1 string = "(())"
var test_2 string = "()()"
var test_3 string = "((("
var test_4 string = "(()(()("
var test_5 string = "))((((("
var test_6 string = "())"
var test_7 string = "))("
var test_8 string = ")))"
var test_9 string = ")())())"

func get_floor(text string) int {
	var i int = 0

	for _, c := range text {
		if c == '(' {
			i += 1
		} else if c == ')' {
			i -= 1
		}
	}

	return i
}

func main() {
	content, err := ioutil.ReadFile(test_file)
	if err != nil {
		log.Fatal(err)
	}

	text := string(content)

	fmt.Println(get_floor(test_1))
	fmt.Println(get_floor(test_2))
	fmt.Println(get_floor(test_3))
	fmt.Println(get_floor(test_4))
	fmt.Println(get_floor(test_5))
	fmt.Println(get_floor(test_6))
	fmt.Println(get_floor(test_7))
	fmt.Println(get_floor(test_8))
	fmt.Println(get_floor(test_9))
	fmt.Println(get_floor(text))
}
