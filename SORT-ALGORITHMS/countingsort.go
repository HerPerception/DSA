/*============================================================================================
The Counting Sort algorithm sorts an array by counting the number of times each value occurs.
It does not compare values like the previous algorithms and only woeks on non negative integers.
It is fast when the number of possible values k is smaller than the number of values n.
==============================================================================================*/

package main

func CountingSort(arr []uint) []uint {
	max := uint(0)
	for _, num := range arr {
		if num > max {
			max = num
		}
	}
	countarr := make([]int, max+1)

	for index := 0; index < len(arr); index++ {
		num := arr[index]
		countarr[num] += 1
	}
	arr = arr[:0]
	for in, num := range countarr {
		for num > 0 {
			arr = append(arr, uint(in))
			num -= 1
		}
	}
	return arr
}

/*============================================================================================
TIME COMPLEXITY
How fast the Counting Sort algorithm runs depends on both the range of possible values k and
the number of values n. In general, time complexity for Counting Sort is O(n+k).
In  a best case scenario, the range of possible values k is very small compared to the number
of values n and Counting Sort has time complexity O(n).
In a worst case scenario, the range of possible values k is very big compared to number of values
n and Counting Sort can have time complexity of O(n²) or even worse.
It is immportant to consider the range of values being sorted before choosing Counting Sort
as your algorithm. Counting Sort only works for NON NEGATIVE values.
If the numbers to be sorted varies a lot in value (large k), and there are few numbers to sort
(small n), the Counting Sort algorithm is not effective.
==============================================================================================*/
