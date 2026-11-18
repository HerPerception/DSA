# THIS IS A HACKERRANK TASK.
# The function is expected to return an INTEGER.
# The function accepts following parameters:
#  1. INTEGER n
#  2. INTEGER p
def pageCount(n, p):
    count = 0
    left = 1
    right = n if n % 2 == 0 else n - 1
    
    while left < right:
        if p <= left or p >= right:
            break
        count += 1
        left += 2
        right -= 2
        
    return count
