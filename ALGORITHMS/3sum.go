package main

import "sort"
func threeSum(nums []int) [][]int {
    sort.Ints(nums)
    result := [][]int{}
    for i := range nums{
        if nums[i] > 0 {
            break
        }
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        seen := make(map[int]int)
        j := i + 1
        for j < len(nums){
            value := -(nums[i] + nums[j])
            if _, exists := seen[value]; exists {
                   result = append(result, []int{nums[i], value, nums[j]})
                   for j+1 < len(nums) && nums[j] == nums[j+1] {
                    j++
                }
            }
             seen[nums[j]] = 1
             j++
        }
    }
    return result
}
