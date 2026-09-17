package main

import "fmt"

func diagonalDifference(arr [][]int32) int32 {
	// Write your code here
	a := int32(0)
	b := int32(0)
	index := 0
	length := 0
	for _, each_arr := range arr {
		length = len(each_arr)
		//fmt.Println(each_arr)
		a += each_arr[index]
		//fmt.Println(each_arr[index])
		index += 1
	}
	length -= 1
	for _, each_arr := range arr {
		//fmt.Println(each_arr)
		b += each_arr[length]
		//fmt.Println(each_arr[length])
		length -= 1
	}
	num := a - b
	if num < 0 {
		num *= -1
	}
	return num
}

func main() {
	fmt.Println(diagonalDifference([][]int32{{11, 2, 4}, {4, 5, 6}, {10, 8, -12}}))
}
