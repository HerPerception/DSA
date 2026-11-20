# Corrected implementation. 199ms runtime and  42.98MB memory use.
from typing import List

class Solution:
    def maxSelectedElements(self, nums: List[int]) -> int:
        nums.sort()
        dp = {}
        
        for v in nums:
            # Update v + 1 first so it uses the previous state of dp[v]
            dp[v + 1] = dp.get(v, 0) + 1
            dp[v] = dp.get(v - 1, 0) + 1
            
        return max(dp.values())
