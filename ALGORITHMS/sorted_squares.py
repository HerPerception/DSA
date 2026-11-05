# Second implementation. Logic was AI assisted. 12ms runtime and 20.99MB memory use.
class Solution:
    def sortedSquares(self, nums: list[int]) -> list[int]:
        left, right = 0, len(nums)-1
        result = [0] * len(nums)
        idx = len(nums)-1
        while left <= right:
            squaredLeft = nums[left] * nums[left]
            squaredRight = nums[right] * nums[right]
            if squaredLeft >= squaredRight:
                result[idx] = squaredLeft
                idx -= 1
                left += 1
            elif squaredRight >= squaredLeft:
                result[idx] = squaredRight
                idx -= 1
                right -= 1
            
        return result
