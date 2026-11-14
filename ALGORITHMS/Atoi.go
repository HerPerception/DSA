package main

// First implementation. Yet to pass all test cases.
import (
    "strings"
    )
func myAtoi(s string) int {
    s = strings.TrimSpace(s)
    if s == "" {
        return 0
    }

    i := 0
    sign := 1
    if s[0] == '-' {
        sign = -1
        i++
    } else if s[0] == '+' {
        i++
    }

    num := 0
    for i < len(s) && s[i] >= '0' && s[i] <= '9' {
        num = num*10 + int(s[i]-'0')
        i++
    }
    return num * sign
}
