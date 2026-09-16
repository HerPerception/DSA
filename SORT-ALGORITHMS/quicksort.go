/*============================================================================================
Quicksort is one of the fastest sorting algorithm. It takes an array of values, chooses one of
the values as the 'central' element, and moves the other values so that the lower values are
on the left of the central element, and the higher values are on the right of it.
==============================================================================================*/

package main

import "fmt"

func Partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1
	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i += 1
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}
func Quicksort(arr []int, low, high int) []int {
	if high == 0 {
		high = len(arr) - 1
	}
	if low < high {
		fmt.Println(arr)
		pivotIndex := Partition(arr, low, high)
		Quicksort(arr, low, pivotIndex-1)
		Quicksort(arr, pivotIndex+1, high)
	}
	return arr
}

/*============================================================================================
TIME COMPLEXITY
The worst case scenario for Quicksortis O(n²). This is when the pivot or central element is 
either the highest or lowest value in every sub-array, which  leads to a lot of recursive calls.
But on average, the time complexity for Quicksort is actually just O(n log n), which is a lot
better than for the previous sorting algorithms we have looked at.
==============================================================================================*/
