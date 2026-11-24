# First implementation. Assisted by AI. 9ms runtime and 19.06MB memory use.
import math
from typing import List

class Solution:
    def maxLength(self, nums: List[int]) -> int:
        max_len = 0
        n = len(nums)
        
        for i in range(n):
            prod = nums[i]
            g = nums[i]
            l = nums[i]
            
            for j in range(i, n):
                if prod > 25200:
                    break
                
                if j > i:
                    prod *= nums[j]
                    g = math.gcd(g, nums[j])
                    l = math.lcm(l, nums[j])
                
                if prod == l * g:
                    max_len = max(max_len, j - i + 1)
                    
        return max_len
