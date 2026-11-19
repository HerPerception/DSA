package main

// 15ms runtime and 11.02MB memory use.
func minimizeSum(nums []int) int {
        sort.Ints(nums)
        gap_1 := nums[len(nums)-1] - nums[2]
        gap_2 := nums[len(nums)-2] - nums[1]
        gap_3 := nums[len(nums)-3] - nums[0]
        return min(gap_1, gap_2, gap_3)
}
