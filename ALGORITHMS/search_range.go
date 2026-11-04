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


// Second implementation. AI assisted. 0ms runtime and 6.33MB memory use. Yet to understand the logic.
func binarySearchInsertion(nums []int, target int) int {
    low := 0
    high := len(nums) - 1
    
    for low <= high {
        mid := low + (high-low)/2 // Fixed the Go syntax here
        
        if nums[mid] == target {
            // Treat it as found, but push high left to find the absolute earliest spot
            high = mid - 1 
        } else if nums[mid] < target {
            low = mid + 1
        } else {
            high = mid - 1
        }
    }
    // When the loop ends, 'low' is the exact index where the target starts/belongs!
    return low
}

func searchRange(nums []int, target int) []int {
    // 1. Find where target starts
    start := binarySearchInsertion(nums, target)
    
    // Check if target is actually in the array at that start spot
    if start == len(nums) || nums[start] != target {
        return []int{-1, -1}
    }
    
    // 2. Find where the NEXT number starts, then step back 1 spot
    end := binarySearchInsertion(nums, target+1) - 1
    
    return []int{start, end}
}
