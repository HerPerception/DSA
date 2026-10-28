package main

// First implementation. 1ms runtime and 4.03MB memory use.
func sumOfMultiples(n int) int {
    total := 0
    for i := 1; i <= n; i++ {
        if i%3 == 0 || i%5 == 0 || i%7 == 0 {
            total += i
        }
    }
    return total
}
