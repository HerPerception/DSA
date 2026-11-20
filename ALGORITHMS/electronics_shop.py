#THIS IS A HACKERRANK TASK.
def getMoneySpent(keyboards, drives, b):
    i = 0
    most_expensive = 0
    while i < len(keyboards):
        j = 0
        while j < len(drives):
            price = keyboards[i] + drives[j]
            if price <= b and price > most_expensive:
                most_expensive = price
            j += 1
        i += 1
    if most_expensive > 0:
        return most_expensive
    else:
        return -1
