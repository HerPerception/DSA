/*============================================================================================
Insertion Sort uses one part of the array to hold sorted values, and the other part of the
array to hold values that are not sorted yet. The algorithm takes one value from the unsorted
part of the array and puts it in the right place in the sorted part of the array, until the
array is fully sorted.
==============================================================================================*/

package main

import "fmt"

func InsertionSort() {
	myArray := []int64{64, 34, 25, 12, 22, 11, 90, 5, 100, 89, 78}
	n := len(myArray)
	for i := 1; i < n; i++ {
		insertIndex := i
		currVal := myArray[i]
		fmt.Println("current value", currVal)
		myArray = append(myArray[:i], myArray[i+1:]...)
		for j := i - 1; j >= 0; j-- {
			if myArray[j] > currVal {
				insertIndex = j
				fmt.Println("j index", myArray[j])
			} else {
				break
			}
		}
		myArray = append(myArray[:insertIndex], append([]int64{currVal}, myArray[insertIndex:]...)...)
		fmt.Println(myArray)
	}
	fmt.Println(myArray)
}
