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
        
# Second implementation. Suggested by AI. 0ms runtime and 19.27MB memory use.
    def checkString2(self, s: str) -> bool:
        return "ba" not in s

# Third implementation. Suggested by AI. 1ms runtime and 18.88MB memory use
    def checkString3(self, s: str) -> bool:
        for i in range(len(s)-1):
            if s[i] == 'b' and s[i+1] == 'a':
                return False

        return True
