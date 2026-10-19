package main

// First implementation. Fully mine. Did not pass all the test cases.
func searchInsert(nums []int, target int) int {
	for i, num := range nums {
		if num == target {
			return i
		} else if i < len(nums)-1 && num < target && nums[i+1] > target {
			return target - num
		}
	}
	return len(nums)
}

// Second implementation. AI assisted.
func searchInsert2(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid // Target found
		} else if nums[mid] < target {
			left = mid + 1 // Search in the right half
		} else {
			right = mid - 1 // Search in the left half
		}
	}

	return left
}
