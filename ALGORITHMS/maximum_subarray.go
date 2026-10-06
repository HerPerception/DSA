package main

//First implemention
func maxSubArray(nums []int) int {
    currentSum := 0
    highestSum := nums[0]
    for _, num := range nums {
        currentSum += num
        if currentSum > highestSum {
            highestSum = currentSum
        }
        if currentSum < 0 {
            currentSum = 0
        }
    }
    return highestSum
}

//Second implementation. Assisted.
func maxSubArray(nums []int) int {
    current, best := nums[0], nums[0]

    for _, num := range nums[1:] {
        current = max(num, current+num)
        best = max(best, current)
    }

    return best
}
