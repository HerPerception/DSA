# First implementation. AI assisted with the logic and debugging. 1ms runtime and 19.09MB memory use.
class Solution:
    def triangleType(self, nums: List[int]) -> str:
        a = nums[0]
        b = nums[1]
        c = nums[2]
        if a + b <= c or a + c <= b or b + c <= a:
            return "none"
        if a == b and b == c:
            return "equilateral"
        elif a == b or b == c  or a == c:
            return "isosceles"

        return "scalene"
            
