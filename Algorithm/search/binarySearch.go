package main

import (
	"fmt"
)

// iterative way
// from lowest to highest, in a sorted manner
// if the middle number is the value given return true
// if the value givn is less than the middle number, move towards the startNum
// if the value given is more than the middle number move towards the endNum
func BinarySearch(data []int, val int) bool {

	startNum := 0
	endNum := len(data) - 1

	for startNum <= endNum {
		midNum := startNum + (endNum-startNum)/2

		if data[midNum] == val {
			return true
		} else if midNum < val {
			startNum++
		} else {
			endNum--
		}
	}
	return false
}

// if the data in row & col equals value, return true
// if its greater than value move to previous col
// else move to next row
func FindElement2DArray(data [][]int, row, col, val int) bool {
	initialRow := 0
	column := col - 1

	for initialRow < row && column >= 0 {
		if data[row][col] == val {
			return true
		} else if data[row][col] > val {
			col--
		} else {
			row++
		}
	}
	return false
}

func FirstNumberToBeRepeated(data []int) int {

	for i := 0; i < len(data); i++ {
		for j := i + 1; j < len(data); j++ {
			if data[i] == data[j] {
				return data[i]
			}
		}
	}
	return 0
}

func main() {

	s := []int{1, 2, 3, 3, 3, 4, 5, 6, 7, 7, 7, 8, 9}

	fmt.Println(BinarySearch(s, 2))
	fmt.Println(FirstNumberToBeRepeated(s))

}
