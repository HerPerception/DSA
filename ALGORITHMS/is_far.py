# First implementation. 42ms runtime and 19.21MB memory use.
class Solution:
    def findTheDistanceValue(self, arr1: list[int], arr2: list[int], d: int) -> int:
        count = 0
        for num in arr1:
            far = True
            for j in range(len(arr2)):
                dis = max(num, arr2[j]) - min(num, arr2[j])
                if dis <= d:
                    far = False
                    break
    
            if far:
                count += 1
            
        return count
