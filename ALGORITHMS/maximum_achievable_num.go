package main

// First implementation. 1ms runtime and 4.66MB memory use.
func theMaximumAchievableX(num int, t int) int {
    return num + (2*t)
}
