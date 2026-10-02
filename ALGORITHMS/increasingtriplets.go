package main

import (
	"fmt"
	"math"
	)

func increasingTriplet(nums []int) bool {
	first := math.MaxInt
	second := math.MaxInt
	for i := range nums {
		if nums[i] <= first {
			first = nums[i]
		}
		if nums[i] > first && nums[i] <= second {
			second = nums[i]
		}
		if nums[i] > first && nums[i] > second {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(increasingTriplet([]int{1, 5, 0, 4, 1, 3}))
	fmt.Println(increasingTriplet([]int{20, 100, 10, 12, 5, 13}))
}

// func increasingTriplet(nums []int) bool {
// 	for i := 0; i < len(nums)-2; i++ {
// 		for j := i + 1; j < len(nums)-1; j++ {
// 			for k := j + 1; k < len(nums); k++ {
// 				if nums[i] < nums[j] && nums[j] < nums[k] {
// 					return true
// 				}
// 			}
// 		}
// 	}
// 	return false
// }

// func increasingTriplet(nums []int) bool {
//     first := 0
//     second := 0
//     third := 0
//     for i := 0; i < len(nums)-1; i++ {
//         for j := i+1; j < len(nums)-1; j++ {
//             if nums[i] < nums[j] {
//                 first = nums[i]
//                 second = nums[j]
//             }
//             for k := j+1; k < len(nums); k++ {
//                 if nums[j] < nums[k] {
//                     third = nums[k]
//                 }
//             }
//         }
//     }
//     return first != 0 && second != 0 && third != 0
// }

// func increasingTriplet(nums []int) bool {
//     i := 0
//     j := i+1
//     k := j+1
//     for i < j  && j < k  && k < len(nums){
//         if nums[i] < nums[j] && nums[j] < nums[k] {
//             return true
//         }
//         i++
//         j++
//         k++
//     }
//     return false
// }
