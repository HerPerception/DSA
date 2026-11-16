package main

import (
    "slices"
)

// THIS IS A HACKERRANK TASK.
/*
 * Complete the 'migratoryBirds' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func migratoryBirds(arr []int32) int32 {
    idDict := make(map[int32]int32)
    for _, val := range arr{
        idDict[val]++
    }
    maxVal := int32(0)
    maxArr := []int32{}
    for key, val := range idDict{
        if val > maxVal{
            maxArr = []int32{}
            maxVal = val
            maxArr = append(maxArr, key)
        }else if val == maxVal{
            maxArr = append(maxArr, key)
        }
    }
    return slices.Min(maxArr)
}
