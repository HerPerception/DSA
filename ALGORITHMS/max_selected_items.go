package main

// First implementation. 117ms runtime and 13.94MB memory use.
func maxSelectedElements(nums []int) int {
    sort.Ints(nums)
    dp := make(map[int]int)
        
    for _, v := range nums{
        // Update v + 1 first so it uses the previous state of dp[v]
        dp[v + 1] = dp[v] + 1
        dp[v] = dp[v - 1] + 1
    } 
    maxVal := 0
    for _, v := range dp {
        if v > maxVal {
            maxVal = v
        }
    }
    return maxVal
}
