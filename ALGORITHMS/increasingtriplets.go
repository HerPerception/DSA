package main

import "fmt"

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
