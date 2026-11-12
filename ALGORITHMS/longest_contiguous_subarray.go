package main

// First implementation. 2ms runtime and 5.05MB memory use.
func longestMonotonicSubarray(nums []int) int {

	best := 1
	up_streak := 1
	down_streak := 1
	for i := 0; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			up_streak += 1
			down_streak = 1
		} else if nums[i] < nums[i-1] {
			down_streak += 1
			up_streak = 1
		} else if nums[i] == nums[i-1] {
			down_streak = 1
			up_streak = 1
		}
		if best < up_streak || best < down_streak {
			best = max(up_streak, down_streak)
		}

	}
	return best

}
