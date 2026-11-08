package main

// First implementation. My own logic. 0ms runtime and 3.78MB memory use.
func countDigits(num int) int {
    val := num
    count := 0
    for val > 0 {
        lastVal := val % 10
        if num%lastVal == 0 {
            count++
        }
        val /= 10
    }
    return count
}
