package main

import "sort"

func firstMissingPositive(nums []int) int {
    sort.Ints(nums)
    number := 1
    for i := range nums {
        if nums[i] < 0 {
            continue 
        } else if nums[i] == number {
            number += 1
        } else if nums[i] > number {
            return number 
        }
    }
    return number 
}
