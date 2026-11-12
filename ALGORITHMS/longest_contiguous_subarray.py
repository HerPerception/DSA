# First implementtion. Assisted by AI. 1ms runtime and 19.15MB memory use.
class Solution:
    def longestMonotonicSubarray(self, nums: List[int]) -> int:
        best = 1
        up_streak = 1
        down_streak = 1
        for i in range(1, len(nums)):
            if nums[i] > nums[i-1]:
                up_streak += 1
                down_streak = 1
            elif nums[i] < nums[i-1]:
                down_streak += 1
                up_streak = 1
            elif nums[i] == nums[i-1]:
                down_streak = 1
                up_streak = 1
            if best < up_streak or best < down_streak:
                best = max(up_streak, down_streak)
        return best
