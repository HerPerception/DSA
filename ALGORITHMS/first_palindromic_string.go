package main

import "strings"

// First implementation. My very own logic. 6ms runtime and 9.2MB memory use.
func firstPalindrome(words []string) string {
    for _, word := range words {
        var revWord strings.Builder
        for i := len(word)-1; i >= 0; i-- {
            revWord.WriteByte(word[i])
        }
        if revWord.String() == word {
            return word
        }
    }
    return ""
}
