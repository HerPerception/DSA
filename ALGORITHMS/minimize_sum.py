# First implementation. Did mnot pass all the test cases.
class Solution:
    def minimizeSum(self, nums: List[int]) -> int:
        nums.sort()
        nums[0], nums[1] = nums[-2], nums[-2]
        return nums[len(nums)-1] - nums[2]
