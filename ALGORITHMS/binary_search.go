package main

// Implemented Binary search after a long time. 0ms runtime and 10.92MB memory use.
func search(nums []int, target int) int {
    left, right := 0, len(nums)-1
    for left <= right{
        mid := (left+right)/2
        if nums[mid] == target {
            return mid
        } else if nums[mid] < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return -1
}
