/*================================================================================================================================
The SELECTON SORT algorithm finds the lowest value in an array and moves it to the front of  the array. It lloks through the array
again and again, moving the next lowest values to the front, until the array is sorted.
================================================================================================================================*/

package main

import "fmt"

func SelectionSort() {
	fmt.Println()
	/*============================================================================================================================
	Here, we swap values instead of shifting thet other values backward.
	==============================================================================================================================*/
	arrayNum := []int{5, 8, 4, 10, 6, 3, 11, 9, 2, 12, 7, 1}
	n := len(arrayNum)
	times := 0
	for times < n-1 {
		minIndex := times
		for in := times + 1; in < n; in++ {
			if arrayNum[in] < arrayNum[minIndex] {
				minIndex = in
			}
		}
		temp := arrayNum[times]
		minVal := arrayNum[minIndex]
		arrayNum[times] = minVal
		arrayNum[minIndex] = temp
		fmt.Println(arrayNum)
		times++
	}

	/*============================================================================================================================
	Here, we shift the other values forward in index position (that is higher index) after bringing the lowest value in front. 
	Because you are bringing the lowest value forward to the front of the line, the other values must shift forward (to the right)
	 to make room for it.
	 * The lowest value moves backward (to a lower index / the front).
	 * The other values shift forward (to a higher index / the back)
	=============================================================================================================================*/
	my_array := []int{64, 34, 25, 5, 22, 11, 90, 12}
	n = len(my_array)
	for i := 0; i < n-1; i++ {
		min_index := i
		for j := i + 1; j < n; j++ {
			if my_array[j] < my_array[min_index] {
				min_index = j
			}
		}
		// 1. POP: Save the value and remove it from the slice
		min_value := my_array[min_index]
		my_array = append(my_array[:min_index], my_array[min_index+1:]...)
		// 2. INSERT: Put min_value at index 'i' by rebuilding the slice
		// We make a temporary slice to hold the pieces safely
		temp := append([]int{min_value}, my_array[i:]...)
		my_array = append(my_array[:i], temp...)
	}
	fmt.Println(my_array)
}

/*=================================================================================================================================
The first implementation of the SELECTION SORT algorithm (swap) is the improved version of the second where you always have 
to shift higher values to the back, instead you just swap the values at both indexes.

TIME COMPLEXITY OF SELECTION SORT
This algorithm sorts an array of n values. On average, about n/2 elements are compared to find the lowest value in each loop.
Selection sort must run the loop to find the lowest value approximately  n times.
Time complexity is O(n²). The runtime is the same as for Bubble Sort: The runtime increases very fast when the size of the array
is increased. 
The most significant difference from Bubble sort that can be noticed here is that best and worts case is actually almost the same
for Selection Sort (O(n²)), but for Bubble Sort the best case runtime is O(n).
==================================================================================================================================*/
