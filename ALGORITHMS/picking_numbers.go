/*==================================================================
THIS IS A HACKERRANK TASK.
====================================================================*/
package main

import (
	"slices" // Use "sort" and sort.Ints(a) if using Go < 1.21
)

func pickingNumbers(a []int32) int32 {
	// Step 1: Sort the array from smallest to largest
	slices.Sort(a)

	// Step 2: Create tracking variables
	var start int32 = 0
	var maxLength int32 = 0

	// Step 3: Loop through the array with the 'end' pointer
	for end := int32(0); end < int32(len(a)); end++ {
		// Shrink the group if the difference is greater than 1
		for a[end]-a[start] > 1 {
			start++
		}

		// Calculate current group size and update record
		currentLength := end - start + 1
		if currentLength > maxLength {
			maxLength = currentLength
		}
	}

	// Step 4: Return the maximum length found
	return maxLength
}
