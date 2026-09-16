package main

import "fmt"

/*
===========================================================================================================================
Implementation using Recursion.
Recursion is when a function calls itself.
=============================================================================================================================
*/
func RecursiveFibonacci(prev1, prev2, count int) int {
	fmt.Println(prev2)
	newFibo := prev1 + prev2
	prev2 = prev1
	prev1 = newFibo
	count += 1
	if count == 19 {
		return prev2
	}
	newFibo = RecursiveFibonacci(prev1, prev2, count)
	return newFibo
}

/*
===========================================================================================================================
Finding the nth number using recursion.
The formula for finding the Fibonacci number is: F(n) = F(n-1) + F(n-2)
This formula uses a zero-based index, meaning that to generate the 20th Fibonacci number, we must write F(19).
=============================================================================================================================
*/

func FibonacciN(n int) int {
	if n <= 1 {
		return n
	}
	return FibonacciN(n-1) + FibonacciN(n-2)
}

/*
===========================================================================================================================
This recursive method calls itself two times whould would make a huge difference in how the program will actually run
on our computer. The number of calculations will explode when we increase the Fibonacci number we want. The number of
function calls would double every time we increase the Fibonacci number we want by one.

IN CONCLUSION, an algorithm can be implemented in differrent ways and in different programming languages.
Recursion and loops are two different programming techniques that can be used to implement algorithms.
=============================================================================================================================
*/
