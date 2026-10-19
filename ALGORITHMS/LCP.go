package main

import "fmt"

func Run(){
	fmt.Println(longestCommonPrefix([]string{"ab", "a"}))
	fmt.Println(longestCommonPrefix([]string{"flower","flow","flight"}))
	fmt.Println(longestCommonPrefix([]string{"dog","racecar","car"}))
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