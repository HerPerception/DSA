package main

func strStr(haystack string, needle string) int {
    for i, ch := range haystack {
        if i + len(needle) <= len(haystack) && ch == rune(needle[0]) && haystack[i:i + len(needle)] == needle {
            return i
        }
    }
    return -1
}
