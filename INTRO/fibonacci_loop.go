/*
=========================================================================================================================
To generate a fibonacci number, we need to add the two previous numbers.
=========================================================================================================================
*/
package main

import (
	"fmt"
)

/*
=========================================================================================================================
Implementation using a For Loop.
=========================================================================================================================
*/
func main() {
	prev1 := 0
	prev2 := 1
	fibonacci := 0
	for range 20 {
		fmt.Println(prev1)
		fibonacci = prev1 + prev2
		prev1 = prev2
		prev2 = fibonacci
	}
	fmt.Println()
	fmt.Println(RecursiveFibonacci(1, 0, 0))
	fmt.Println()
	fmt.Println(FibonacciN(19))
}
