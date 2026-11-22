# First implementation. 0ms runtime and 19.03MB memory use.
class Solution:
    def checkString(self, s: str) -> bool:
        b_seen = False
        for ch in s:
            if ch == 'b':
                b_seen = True
            if ch == 'a' and b_seen:
                return False

        return True
