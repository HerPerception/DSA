package main

/*
The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER p
 */

func pageCount(n int32, p int32) int32 {
    var count int32 = 0
    var left int32 = 1
    right := n
    if n%2 != 0 {
        right = n - 1
    }

    for left < right {
        if p <= left || p >= right {
            break
        }
        count++
        left += 2
        right -= 2
    }
    return count
}
