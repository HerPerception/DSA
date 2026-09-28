package main

func birthdayCakeCandles(candles []int32) int32 {
    // Write your code here
    tallest := int64(0)
    count := int32(0)
    for _, num := range candles {
        if int64(num) > tallest {
            tallest = int64(num)
            count = 1
        } else if int64(num) == tallest {
            count++
        }
    }
    return count
}
