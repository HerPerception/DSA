package main

import "fmt"

func miniMaxSum(arr []int32) {
    total := int64(0) 
    for _, num := range arr {
        total += int64(num)
    }
    min_total := total
    max_total := int64(0)
    for _, num := range arr{
        temp := total
        temp -= int64(num)
        if min_total > temp {
            min_total = temp
        } 
        if max_total < temp {
            max_total = temp
        }
    }
    fmt.Println(min_total, max_total)
}
