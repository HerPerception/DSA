package main

// First implementation. 1ms runtime and 7.32MB memory use.
func minChanges(s string) int {
    count, i := 0, 0
    for i < len(s)-1 {
        if s[i] != s[i+1] {
            count++
        }
        i+= 2
    }
    return count
}
