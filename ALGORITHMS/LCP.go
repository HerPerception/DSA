package main

import "fmt"

func main(){
	fmt.Println(longestCommonPrefix([]string{"ab", "a"}))
	fmt.Println(longestCommonPrefix2([]string{"ab", "a"}))
	fmt.Println(longestCommonPrefix2([]string{"flower","flow","flight"}))
	fmt.Println(longestCommonPrefix2([]string{"dog","racecar","car"}))
}

// First implementation after running from this task for months. Sometimes, what you're afraid of 
// is not the prooblem, your mind is. 0ms runtime and 4.40MB memory use. My very own logic.
func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }
    if len(strs) == 1 {
        return strs[0]
    }
   
    LCP := strs[0]
    for i := 1; i < len(strs); i++ {
        word := strs[i]
        if word == "" {
            return ""
        }
        j := 0
        for j < len(word) && j < len(LCP) {
            if word[j] == LCP[j] {
                j++
            } else {
                break
            }
        }
        LCP = LCP[:j]
    }
    return LCP
}

//Second implementation. Mrutyunjaya's logic (he's on leetcode. He used Python and I decided to try it in Golang.)
//0ms runtime and 4.31MB memory use.
func longestCommonPrefix2(strs []string) string {
    prefix := strs[0]
    prefix_length := len(prefix)
    for i := 1; i < len(strs); i++{
        word := strs[i]
        if prefix_length > len(word) {
            prefix_length = len(word)
        }
		prefix = prefix[0:prefix_length]
        for prefix != word[0:prefix_length] {
            prefix_length -= 1
            if prefix_length == 0 {
                return ""
            }
			prefix = prefix[0:prefix_length]
        }
		
    }
    return prefix
}