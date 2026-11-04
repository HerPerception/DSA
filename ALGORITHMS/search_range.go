package main

// First implementation. Did not pass the restriction. (The implementation uses a linear scan (for loop over nums) which results in O(n) time complexity instead of the required O(log n).)
func searchRange(nums []int, target int) []int {
    target_index := []int{}
        for i := range nums {
            if nums[i] == target && len(target_index) == 0 {
                target_index = append(target_index, i)
            }else if nums[i] == target && nums[i+1] != target {
                target_index = append(target_index, i)
                break
            }
    }
    if len(target_index) == 0 {
        return []int{-1, -1}
    }
    return target_index
}
