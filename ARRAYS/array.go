/*=============================================================================
An array is a data structure used to hold/store multiple elements.
An algorithm can be used too look through an array to find the lowest or
the highest value.
===============================================================================*/

package main

import "fmt"

func main() {
	my_array := []int{1, 2, 7, 2, 4, 6, 7, 8, 9}
	fmt.Println(my_array[0])
	// FINDING THE LOWEST AND HIGHEST VALUE IN AN ARRAY
	lwst := my_array[0]
	for _, num := range my_array {
		if num < lwst {
			lwst = num
		}
	}
	fmt.Println(lwst)
}

/*=============================================================================
Arrays are indexed, each element in an array has an index that says where the
element is located in the array. Go uses zero-based indexing.

PSEUDOCODE is a description of what a program does, using a language that is 
something between the human language and a programming language.

TIME COMPLEXITY: how much time an algorithm takes to run relative to the size 
of the dataset.
We say that an operation takes constant time if it takes the same amount of time
regardless of the amount of data (n) the algorithm is processing.
Big O Notation is used to find the worst case time complexity for an algorithm
===============================================================================*/
