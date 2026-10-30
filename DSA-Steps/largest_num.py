def find_largest(nums: list[int]):
    if nums == []:
        return None
    highest = 0
    for _, val in enumerate(nums):
        if highest < val:
            highest = val
    return highest

print(find_largest([3, 4, 1, 2, 7, 6, 4]))