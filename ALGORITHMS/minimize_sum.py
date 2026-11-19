# First implementation. Did mnot pass all the test cases.
class Solution:
    def minimizeSum(self, nums: List[int]) -> int:
        nums.sort()
        nums[0], nums[1] = nums[-2], nums[-2]
        return nums[len(nums)-1] - nums[2]
# Second implementation. 34ms runtime and 31.41MB memory use.
    def minimizeSum2(self, nums: List[int]) -> int:
        nums.sort()
        gap_1 = nums[len(nums)-1] - nums[2]
        gap_2 = nums[len(nums)-2] - nums[1]
        gap_3 = nums[len(nums)-3] - nums[0]
        return min(gap_1, gap_2, gap_3)
