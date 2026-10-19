package main

import "math"

//First implementation. Passed 990 of 994 test cases
func divide(dividend int, divisor int) int {
    neg := false
if divisor < 0 {
    divisor = -divisor  
    neg = !neg
}

if dividend < 0 {
    dividend = -dividend  
    neg = !neg
}

    count := 0
    for dividend > 0 {
        if dividend < divisor {
            break
        }
        dividend -= divisor
        count += 1
    }
    if neg {
        count = -count
    }
    return count
}

//Second implementation. AI assissted. Passed all test cases
func divide2(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	neg := false
	if divisor < 0 {
		divisor = int(-int64(divisor))
		neg = !neg
	}
	if dividend < 0 {
		dividend = int(-int64(dividend))
		neg = !neg
	}
    
	count := 0
    
	for dividend >= divisor {
		tempDivisor := divisor
		multiple := 1
		
		for dividend >= (tempDivisor << 1) {
			tempDivisor <<= 1
			multiple <<= 1
		}
		
		dividend -= tempDivisor
		count += multiple
	}
    
	if neg {
		count = -count
	}
	return count
}
