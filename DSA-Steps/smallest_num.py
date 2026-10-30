def find_smallest(nums: list[int]):
    if nums == []:
        return None
    smallest = nums[0]
    for _, val in enumerate(nums):
        if smallest > val:
            smallest = val
    return smallest

print(find_smallest([3, 4, 1, 2, 7, 6, 4]))
print(find_smallest([-3, -10, -1, -6]))