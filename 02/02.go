package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var input_file string = "input.txt"

func getDimensions(textLine string) (int, int, int) {
	splitLine := strings.Split(textLine, "x")

	l, _ := strconv.Atoi(splitLine[0])
	w, _ := strconv.Atoi(splitLine[1])
	h, _ := strconv.Atoi(splitLine[2])

	return l, w, h
}

func getMin(arr []int) int {
	var min = arr[0]

	for _, i := range arr {
		if i < min {
			min = i
		}
	}

	return min
}

func removeValFromArray(arr []int, val int) []int {
	var index int = 0

	for i, check := range arr {
		if check == val {
			index = i
		}
	}

	arr[index] = arr[len(arr)-1]
	return arr[:len(arr)-1]
}

func calculateWrappingPaper(l int, w int, h int) int {
	lw := l * w
	lh := l * h
	wh := w * h

	areaArray := []int{lw, lh, wh}

	min := getMin(areaArray[:])

	total := (2 * lw) + (2 * lh) + (2 * wh) + min

	return total
}

func calculateRibbon(l int, w int, h int) int {
	var lengthArray []int = []int{l, w, h}

	min_1 := getMin(lengthArray)

	lengthArray = removeValFromArray(lengthArray[:], min_1)

	min_2 := getMin(lengthArray)

	total := (min_1 * 2) + (min_2 * 2) + (l * w * h)

	return total
}

func main() {
	content, err := ioutil.ReadFile(input_file)
	if err != nil {
		log.Fatal(err)
	}

	dimensions := string(content)

	var wrappingPaperTotal int = 0
	var ribbonTotal int = 0

	scanner := bufio.NewScanner(strings.NewReader(dimensions))
	for scanner.Scan() {
		textLine := scanner.Text()
		l, w, h := getDimensions(textLine)
		wrappingPaperTotal += calculateWrappingPaper(l, w, h)
		ribbonTotal += calculateRibbon(l, w, h)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("error occurred: %v\n", err)
	}

	fmt.Println(wrappingPaperTotal, ribbonTotal)
}
