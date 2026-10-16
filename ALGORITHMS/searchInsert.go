package main

//Second implementation. AI assisted.
func searchInsert(nums []int, target int) int {
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
