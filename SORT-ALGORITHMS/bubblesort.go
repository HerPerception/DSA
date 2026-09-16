/*=============================================================================================
BUBBLE SORT is an algorithm that sorts an array from lowest to highest value. It makes the
highest values bubble up.
===============================================================================================*/

package main

import "fmt"

func BubbleSort() {
	arrayNum := []int{8, 7, 4, 6, 5, 9, 0, 3, 10, 2, 1}
	times := 0
	/*=======================================================================================
	1. My original implementation
	==========================================================================================*/
	for times < len(arrayNum) {
		for in := 0; in < len(arrayNum); in++ {
			if in < len(arrayNum)-1 && arrayNum[in] > arrayNum[in+1] {
				temp := arrayNum[in]
				arrayNum[in] = arrayNum[in+1]
				arrayNum[in+1] = temp
			}
		}
		times++
		fmt.Println(arrayNum)
	}
	fmt.Println()
	/*========================================================================================
	 2. Using the lesson learned on DSA
	==========================================================================================*/
	arrayNum = []int{8, 7, 4, 6, 5, 9, 0, 3, 10, 2, 1}
	n := len(arrayNum)
	for count := range arrayNum {
		for in := 0; in < (n - count - 1); in++ {
			if in < len(arrayNum)-1 && arrayNum[in] > arrayNum[in+1] {
				temp := arrayNum[in]
				arrayNum[in] = arrayNum[in+1]
				arrayNum[in+1] = temp
			}
		}
		count++
		fmt.Println(arrayNum)
	}
	fmt.Println()

	/*========================================================================================
	3. Improving the Bubble Sort algorithm so it stops sorting if the algorithm
	goes through the array one time without swapping any value.
	==========================================================================================*/
	arrayNum = []int{8, 7, 4, 6, 5, 9, 0, 3, 10, 2, 1}
	for count := range arrayNum {
		swapped := false
		for in := 0; in < (n - count - 1); in++ {
			if in < len(arrayNum)-1 && arrayNum[in] > arrayNum[in+1] {
				temp := arrayNum[in]
				arrayNum[in] = arrayNum[in+1]
				arrayNum[in+1] = temp
				swapped = true
			}
		}
		if !swapped {
			break
		}
		count++
		fmt.Println(arrayNum)
	}

	/*========================================================================================
	The processes of numbers 1 and 2 runs 11 times while the third implementation runs until no
	number remains to be swapped - 9 times.

	TIME COMPLEXITY
	The time complexity for this algorithm is O(n²): for an array of n values, there must be n 
	comparisons done and looped through again and again n times. This means for each of those 
	items, it has to look through all 10 items again. n stands for the number of items in your 
	list. Because the algorithm multiplies the number of items by itself, the maths is n * n, 
	written as n². The run time increases really fast when the size of the array is increased.
	==========================================================================================*/
	
}
