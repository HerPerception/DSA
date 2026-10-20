package main

//First implementation. AI assisted. 27ms runtime and 8.22MB memory use.
func lengthOfLongestSubstring(s string) int {
    left, right := 0, 0
    hashMap := make(map[rune]int)
    maxLength := 0
    for right < len(s) {
        char := rune(s[right])
        if val, exists := hashMap[char]; exists {
            if val >= left {
                left = val + 1
            }
        }
        hashMap[char] = right
        length := right - left + 1
        if maxLength < length {
            maxLength = length
        }
        right++
    }
    return maxLength
}
