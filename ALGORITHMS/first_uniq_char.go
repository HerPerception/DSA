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

//Second implementation. AI assisted. 0ms runtime, 8.03MB memory.

func firstUniqChar(s string) int {
    var counts [26]int
    
    for i := 0; i < len(s); i++ {
        counts[s[i]-'a']++
    }
    
    for i := 0; i < len(s); i++ {
        if counts[s[i]-'a'] == 1 {
            return i
        }
    }
    
    return -1
}
