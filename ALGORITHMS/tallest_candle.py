def birthdayCakeCandles(candles):
    tallest = 0
    count = 0
    for num in candles:
        if num > tallest:
            tallest = num
            count = 1
        elif num == tallest:
            count += 1
    return count
    
