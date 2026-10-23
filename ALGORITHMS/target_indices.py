class Solution:
    def targetIndices(self, nums: list[int], target: int) -> list[int]:
        nums = sorted(nums)
        target_list = []
        for index, num in enumerate(nums):
            if num == target:
                target_list.append(index)
        return target_list
