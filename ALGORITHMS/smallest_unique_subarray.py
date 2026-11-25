# AI. Yet to understand it.
from collections import Counter

def smallestUniqueSubarray(nums: list[int]) -> int:
    def has_unique(length: int) -> bool:
        b1, b2 = 1000003, 1000033
        mod1, mod2 = 10**9 + 7, 10**9 + 9
        
        h1 = h2 = 0
        pow1 = pow2 = 1
        
        for i in range(length):
            val = nums[i] + 1
            h1 = (h1 * b1 + val) % mod1
            h2 = (h2 * b2 + val) % mod2
            if i > 0:
                pow1 = (pow1 * b1) % mod1
                pow2 = (pow2 * b2) % mod2
                
        counts = Counter([(h1, h2)])
        
        for i in range(length, len(nums)):
            out_val = nums[i - length] + 1
            in_val = nums[i] + 1
            
            h1 = ((h1 - out_val * pow1) * b1 + in_val) % mod1
            h2 = ((h2 - out_val * pow2) * b2 + in_val) % mod2
            
            counts[(h1, h2)] += 1
            
        return 1 in counts.values()

    low, high = 1, len(nums)
    ans = len(nums)
    
    while low <= high:
        mid = low + (high - low) // 2
        if has_unique(mid):
            ans = mid
            high = mid - 1
        else:
            low = mid + 1
            
    return ans
