# AI assisted. First implentation.
class Solution:
    def myAtoi(self, s: str) -> int:
        s = s.strip()
        if not s:
            return 0

        i = 0
        sign = 1
        if s[0] == '-':
            sign = -1
            i += 1
        elif s[0] == '+':
            i += 1

        num = 0
        while i < len(s) and s[i].isdigit():
            digit = int(s[i])

            if num > (2147483647 - digit) // 10:
                return 2147483647 if sign == 1 else -2147483648

            num = num * 10 + digit
            i += 1

        return num * sign
