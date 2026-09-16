package main

import "fmt"

func main() {
	//BubbleSort()
	//fmt.Println()
	//InsertionSort()
	myArray := []int{100, 150, 64, 34, -2, 25, 3, -3, 2, 10, 12, 22, 11, 90, 5}
	low, high := 0, 0
	fmt.Println(Quicksort(myArray, low, high))
	//fmt.Println()
	//SelectionSort()
}
