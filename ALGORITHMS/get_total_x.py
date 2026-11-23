# THIS IS A HACKERRNK TASK.

# Complete the 'getTotalX' function below.

# The function is expected to return an INTEGER.
# The function accepts following parameters:
#  1. INTEGER_ARRAY a
#  2. INTEGER_ARRAY b

def getTotalX(a, b):
    start = max(a)
    end = min(b)
    if start > end:
        return 0
    count = 0
    for x in range(start, end+1):
        if all(x % num == 0 for num in a)and  all(num % x == 0 for num in b):
            count += 1
    return count
    
