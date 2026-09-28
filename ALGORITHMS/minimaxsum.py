def miniMaxSum(arr):
    total = sum(arr)
    max_total = 0
    min_total = total
    for num in arr:
        temp = total
        temp -= num
        if temp < min_total:
            min_total = temp
        if temp > max_total:
            max_total = temp
    print(min_total, max_total)
 
