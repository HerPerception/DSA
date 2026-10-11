package main
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
