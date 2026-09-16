package main

import "fmt"

func main() {
	BubbleSort()
	fmt.Println()
	InsertionSort()
	fmt.Println()
	SelectionSort()
	myArray := []int{100, 150, 64, 34, -2, 25, 3, -3, 2, 10, 12, 22, 11, 90, 5}
	low, high := 0, 0
	fmt.Println(Quicksort(myArray, low, high))
	fmt.Println()
	arr := []uint{7, 8, 9, 7, 2, 3, 4, 2, 3, 1, 4, 1, 6, 7, 6}
	fmt.Println(CountingSort(arr))
	fmt.Println()
}
