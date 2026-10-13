package main

//First Implementation. 23ms runtime and more than 10MB memory use.
func subarraySum(nums []int, k int) int {
    count := 0
    currSum := 0
    targetMap := make(map[int]int)
    targetMap[0] = 1
    for _, num := range nums {
        currSum += num
        need := currSum - k
        if _, exists := targetMap[need]; exists {
            count += targetMap[need]
        }
        targetMap[currSum] = targetMap[currSum] + 1
    }
    return count
}
