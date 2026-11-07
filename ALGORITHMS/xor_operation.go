package main

// First impementation. AI assisted with the logic as well as helping me understand the use of bitwise operator. 0ms runtime and 3.77MB memory use.
func xorOperation(n int, start int) int {
    result := 0
    for i := 0; i < n; i++ {
        currentNum := start + 2 * i
        result ^= currentNum
    }
    return result
}
