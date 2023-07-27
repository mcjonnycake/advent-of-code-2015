package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

const rightArrow rune = '>'
const leftArrow rune = '<'
const upArrow rune = '^'
const downArrow rune = 'v'

var inputPath string = "input.txt"
var testCase1 string = ">"
var testCase2 string = "^>v<"
var testCase3 string = "^v^v^v^v^v"

func getInput(inputPath string) string {
	content, err := ioutil.ReadFile(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}

func buildPointArray(directionInput string) [][]int {
	currentLocation := [][]int{{0, 0}}

	var pointArray [][]int

	firstPoint := make([]int, 2)
	copy(firstPoint, currentLocation[0])

	pointArray = append(pointArray, firstPoint)

	for _, direction := range directionInput {
		coords := currentLocation[0]

		if direction == rightArrow {
			coords[0] = coords[0] + 1
		} else if direction == leftArrow {
			coords[0] = coords[0] - 1
		} else if direction == upArrow {
			coords[1] = coords[1] + 1
		} else if direction == downArrow {
			coords[1] = coords[1] - 1
		}

		newPoint := make([]int, 2)
		copy(newPoint, coords)

		pointArray = addNewPoint(newPoint, pointArray)
	}

	return pointArray
}

func addNewPoint(point []int, expandArray [][]int) [][]int {
	pointInArray := false

	for _, elem := range expandArray {
		elem_x := elem[0]
		elem_y := elem[1]

		if elem_x == point[0] && elem_y == point[1] {
			pointInArray = true
		}
	}

	if pointInArray != true {
		expandArray = append(expandArray, point)
	}

	return expandArray
}

func buildEveryOtherPointArray(directionInput string) ([][]int, [][]int) {
	currentLocation := [][]int{{0, 0}, {0, 0}}

	var pointArray1 [][]int
	var pointArray2 [][]int

	firstPoint := make([]int, 2)
	copy(firstPoint, currentLocation[0])

	pointArray1 = append(pointArray1, firstPoint)
	pointArray2 = append(pointArray2, firstPoint)

	for i, direction := range directionInput {
		coords := currentLocation[i%2]

		if direction == rightArrow {
			coords[0] = coords[0] + 1
		} else if direction == leftArrow {
			coords[0] = coords[0] - 1
		} else if direction == upArrow {
			coords[1] = coords[1] + 1
		} else if direction == downArrow {
			coords[1] = coords[1] - 1
		}

		newPoint := make([]int, 2)
		copy(newPoint, coords)

		if i%2 == 1 {
			pointArray1 = addNewPoint(newPoint, pointArray1)
		} else {
			pointArray2 = addNewPoint(newPoint, pointArray2)
		}
	}

	return pointArray1, pointArray2
}

func combinePointArrays(pointArray1 [][]int, pointArray2 [][]int) [][]int {
	var combinedPointArray [][]int
	var uniqueCombinedPointArray [][]int

	for _, elem := range pointArray1 {
		combinedPointArray = append(combinedPointArray, elem)
	}
	for _, elem := range pointArray2 {
		combinedPointArray = append(combinedPointArray, elem)
	}

	for _, elem := range combinedPointArray {
		uniqueCombinedPointArray = addNewPoint(elem, uniqueCombinedPointArray)
	}

	return uniqueCombinedPointArray
}

func main() {
	input := getInput(inputPath)
	pointArray := buildPointArray(input)
	houseCount := len(pointArray)
	fmt.Println(houseCount)

	pointArray1, pointArray2 := buildEveryOtherPointArray(input)
	uniqueCombinedPointArray := combinePointArrays(pointArray1, pointArray2)
	houseCount = len(uniqueCombinedPointArray)
	fmt.Println(houseCount)
}
