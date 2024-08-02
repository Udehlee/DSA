package main

import "fmt"

// value to be inserted is temporary stored in the holdInsertVal
func insertionSort(data []int, comp func(int, int) bool) {
	var holdInsertVal int
	j := 0

	for i := 1; i < len(data); i++ {
		holdInsertVal = data[i]
		for j := i; j > 0 && comp(data[j-1], holdInsertVal); j-- {
			data[j] = data[j-1]
		}
		data[j] = holdInsertVal
	}
}

//selectionSort puts the largest value at the end of the array
func selectionSort(arr []int) {
	var min int

	for i := 0; i < len(arr)-1; i++ {
		min = i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
}
func main() {

	data := []int{20, 9, 1, 8, 18, 2, 7, 15, 3, 6, 4, 5}
	selectionSort(data)
	fmt.Println(data)

}
