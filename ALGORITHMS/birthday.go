package main

// THIS IS A HACKERRANK TASK.
/* The function is expected to return an INTEGER that counts how many times total equals d.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY s
 *  2. INTEGER d
 *  3. INTEGER m
 */

func birthday(s []int32, d int32, m int32) int32 {
    count := int32(0)
    stop := len(s) - int(m)
    for i := 0; i <= stop; i++ {
        total := 0
        for j := i; j < i+ int(m); j++ {
            total += int(s[j])
        }
        if total == int(d) {
                count += 1
            }
    }
    return count
}
