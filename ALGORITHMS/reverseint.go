package main

import (
  "math"
  "strconv"
)

func reverse(x int) int {
    if x < math.MinInt32 || x > math.MaxInt32 {
		return 0
	}
    neg := false
    if x < 0 {
        neg = true
        x *= -1
    }
    textnum := strconv.Itoa(x)
    numtext := ""
    for i := len(textnum)-1; i>= 0; i-- {
        numtext += string(textnum[i])
    }
    num, err := strconv.Atoi(numtext)
    if err != nil {
        return 0
    }
     if num < math.MinInt32 || num > math.MaxInt32 {
        return 0
    }
    if neg == true { 
        num *= -1
    }
    return num
}
