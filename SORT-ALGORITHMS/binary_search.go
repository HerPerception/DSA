package main

// Finally learned binary search basic after running away for a long time.
nums = [1, 3, 5, 6, 7, 8, 11, 13, 15]
target = 8
left, right = 0, len(nums)-1
while left <= right:
    mid = (left+ right)
    mid = (left+ right)//2
    if nums[mid] == target:
        print(mid)
        break
    elif nums[mid] < target:
        left = mid + 1
    else:
        right = mid - 1


