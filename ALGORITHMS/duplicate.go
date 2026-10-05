package main

func containsDuplicate(nums []int) bool {
    numberMap := make(map[int]int)
    for _, num := range nums {
        if _, exists := numberMap[num]; exists {
            return true
        } else {
            numberMap[num] = num
        }
    }
    return false
}
