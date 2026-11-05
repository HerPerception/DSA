package main

// First implementation. Logic was AI assisted. 0ms runtime and 9.21MB memory use.
func sortedSquares(nums []int) []int {
    left, right := 0, len(nums)-1
    result := make([]int, len(nums))
    idx := len(nums)-1
    for left <= right {
        squaredLeft := nums[left] * nums[left]
        squaredRight := nums[right] * nums[right]
        if squaredLeft >= squaredRight {
            result[idx] = squaredLeft
            idx--
            left++
        } else if squaredRight >= squaredLeft {
            result[idx] = squaredRight
            idx--
            right--
        }
        
    }
    return result
}
