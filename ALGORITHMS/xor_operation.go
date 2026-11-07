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


// The bitwise XOR (Exclusive OR) operator compares two numbers bit by bit (in their binary 0 and 1 formats) and applies one simple rule:
// It returns a 1 if the two bits are different, and a 0 if they are the same.
// Think of it like a strict "either-or" switch:
// • 1 and 0 → 1 (They are different)
// • 0 and 1 → 1 (They are different)
// • 0 and 0 → 0 (They are the same)
// • 1 and 1 → 0 (They are the same)
