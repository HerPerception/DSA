package main

//First implementation. 19ms runtime
func firstUniqChar(s string) int {
    strMap := make(map[rune]int)

    for _, ch := range s {
        if _, exists := strMap[ch]; exists {
            strMap[ch] += 1
        } else {
            strMap[ch] = 1
        }
    }

    for i, ch := range s {
        if strMap[ch] == 1 {
            return i
        }
    }
    return -1
}
