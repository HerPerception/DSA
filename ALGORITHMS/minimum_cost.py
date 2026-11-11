# First implementation. 1ms runtime and 19.24MB memory use.
class Solution:
    def minimumCost(self, cost1: int, cost2: int, costBoth: int, need1: int, need2: int) -> int:
        totalA = (need1 * cost1) + (need2 * cost2)
        overlap = min(need1, need2)
        rem1 = need1 - overlap
        rem2 = need2 - overlap
        rem1 *= cost1
        rem2 *= cost2
        overlap *= costBoth
        totalB = overlap + rem1 + rem2
        totalC = max(need1, need2) * costBoth
        if totalA <= totalB and totalA <= totalC:
            return totalA
        elif totalB <= totalA and totalB <= totalC:
            return totalB

        return totalC
