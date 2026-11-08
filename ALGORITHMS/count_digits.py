# First implemention. Same logic I implemented in Golang, I thought to try it in Python. This task reminded me of a division rule in Python that I had forgotten: 
# Using the '/' sign would give both the whole and the decimal but the '//' sign would only give the whole number and ignore the decimal should the be a decimal. 
#I forgot the rule because the original implementation was in Golang and a rule like that does not exist in the language. The code looked clean and correct and I
#just could not find the bug, until I looked it up with Google and my error was pointed out. 
#0ms runtime and 19.22MB memory use.
class Solution:
    def countDigits(self, num: int) -> int:
        val = num
        count = 0
        while val > 0:
            lastVal = val % 10
            if num%lastVal == 0:
                count += 1
            val //= 10

        return count
            
