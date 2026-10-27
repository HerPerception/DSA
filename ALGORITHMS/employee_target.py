# First implementation. My own logic, debugged with AI assistance. 0ms runtime and 19.1MB
class Solution:
    def numberOfEmployeesWhoMetTarget(self, hours: List[int], target: int) -> int:
        higher = False
        hours.sort()
        for index in range(len(hours)):
            if hours[index] >= target:
                higher = True
                break
        if higher == True:
            return len(hours[index:])
        else:
            return 0
        
