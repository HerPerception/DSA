/*============================================================================================
The Counting Sort algorithm sorts an array by counting the number of times each value occurs.
It does not compare values like the previous algorithms and only woeks on non negative integers.
It is fast when the number of possible values k is smaller than the number of values n.
==============================================================================================*/

package main

import "fmt"

func CountingSort(arr []int) []int {
	countarr := make([]int, len(arr))

	for index := 0; index < len(arr); index++ {
		fmt.Println(arr)
		fmt.Println(countarr)
		num := arr[index]
		if countarr[num] >= 1 {
			countarr[num] += 1
		} else {
			countarr[num] = 1
		}
		fmt.Println(index)
		fmt.Println(len(arr))
		arr = append(arr[:index], arr[index+1:]...)
		index--
	}
	for in, num := range countarr {
		for num > 0 {
			arr = append(arr, in)
		}
	}
	return arr
}

func main() {
	arr := []int{7, 8, 9, 7, 2, 3, 4, 2, 3, 1, 4, 1, 6, 7, 6}
	fmt.Println(CountingSort(arr))
}
