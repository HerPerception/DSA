# THIS IS A HACKERRANK TASK.
def birthday(s, d, m):
    count = 0
    stop = len(s) - m
    for i in range(stop+1):
        total = 0
        for j in range(i, i+m):
            total += s[j]
        
        if total == d:
            count += 1

    return count
