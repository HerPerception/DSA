# THIS IS A HACKERRANK TASK.
#First implementation. Logic is mine but debugging was assisted by AI.
def divisibleSumPairs(n, k, ar):
    count = 0
    for i in range(n-1):
        for j in range(i+1, n):
            if (ar[i] + ar[j])%k == 0:
                count += 1
    return count 
