package main

import (
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	var builder strings.Builder
	for i := 0; i < len(s); i++ {
		c := s[i]
		// keep only a-z and 0-9
		if (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') {
			builder.WriteByte(c)
		}
	}
	text := builder.String()
	left, right := 0, len(text)-1
	for left < right {
		if text[left] != text[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// Second implementation.
func isPalindrome2(s string) (string, bool) {
	s = strings.ToLower(s)
	var builder strings.Builder
	for i := 0; i < len(s); i++ {
		c := s[i]
		// keep only a-z and 0-9
		if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
			builder.WriteByte(c)
		}
	}
	text := builder.String()
	left, right := 0, len(text)-1
	for left < right {
		if text[left] != text[right] {
			return text, false
		}
		left++
		right--
	}
	return text, true
}

// func main() {
//  fmt.Println(diagonalDifference([][]int32{{11, 2, 4}, {4, 5, 6}, {10, 8, -12}}))
//  }
