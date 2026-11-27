# First implementation. 0ms runtime and 19.09MB memory use.
class Solution:
    def canAliceWin(self, nums: List[int]) -> bool:
        single = 0
        double = 0
        for num in nums:
            if num < 10:
                single += num
            else:
                double += num
            
        if single > double or double > single: 
            return True
        
        return False
