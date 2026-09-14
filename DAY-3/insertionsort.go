/*============================================================================================
Insertion Sort uses one part of the array to hold sorted values, and the other part of the
array to hold values that are not sorted yet. The algorithm takes one value from the unsorted
part of the array and puts it in the right place in the sorted part of the array, until the
array is fully sorted.
==============================================================================================*/

package main

import "fmt"

func InsertionSort() {
	myArray := []int{64, 34, 25, 12, 22, 11, 90, 5, 3, 50, 100, 300, 200, 2}
	n := len(myArray)
	for i := 1; i < n; i++ {
		insertIndex := i
		currVal := myArray[i]
		fmt.Println(currVal)
		for j := i - 1; j >= 0; j-- {
			if myArray[j] > currVal {
				insertIndex = j
				myArray[j+1] = myArray[j]
				fmt.Println(myArray)
				myArray[insertIndex] = currVal
				fmt.Println(myArray)
			} else {
				break
			}
		}
		fmt.Println(myArray)
	}
	fmt.Println(myArray)
}

/*============================================================================================
We break out of the inner loop because there's no need comparing values when the correct place
for the currrent value is found.
TIME COMPLEXITY
Insertion Sort sorts an array of n values.
On average, each value must be compared to about n\2 other values to find the correct place 
to insert it.
Insertion Sort must run the loop to insert a value in its correct place approximately n times.
Time complexity is O(n²)
For Insertion Sort, there is a big difference between best, average and worst case scenarios. 
==============================================================================================*/

