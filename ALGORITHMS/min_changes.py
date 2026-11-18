# First implementation. Problem understanding and logic assisted by AI. 25ms runtime and 19.59MB memory use.
class Solution:
    def minChanges(self, s: str) -> int:
        i = 0
        count = 0
        while i < len(s)-1:
            if s[i] != s[i+1]:
                count += 1
            i += 2
        return count
