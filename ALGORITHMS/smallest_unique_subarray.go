package main

// First implementation. Exceeded time limit for some of the test cases.
func smallestUniqueSubarray(nums []int) int {
    for i := 1; i <= len(nums); i++ {
        subMap := make(map[string]int)
        for j := 0; j <= len(nums)-i; j++ {
            subarr := nums[j:j+i]
            substr := fmt.Sprint(subarr)
            subMap[substr]++
        }
        for _, count := range subMap {
            if count == 1 {
            return i
            }
        }
    }
    return len(nums)
}
