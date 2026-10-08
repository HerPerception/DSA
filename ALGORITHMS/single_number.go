package main

func singleNumber(nums []int) int {
    seenMap := make(map[int]int)
    for _, num := range nums {
        seenMap[num]++
    }

    for num, count := range seenMap {
        if count == 1 {
            return num
        }
    }
    return 0 
}
