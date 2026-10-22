package main
//First implementation. My logic. 0ms runtime and 3.93 memory use
func isSubsequence2(s string, t string) bool {
    i := 0
    j := 0
    strText := ""
    for i < len(s) && j < len(t) {
        if s[i] == t[j] {
            strText += string(s[i])
            i++
            j++
        } else if s[i] != t[j] {
            j++
        } 
    }
    if strText == s {
        return true
    }
    return false
}

//Second implementation. My logic. 0ms runtime and 3.94 memory use
func isSubsequence(s string, t string) bool {
    if len(s) == 0 {
        return true
    }
    i := 0
    j := 0
    for i < len(s) && j < len(t) {
        if s[i] == t[j] {
            i++
            j++
        } else if s[i] != t[j] {
            j++
        } 
    }
   return i == len(s)
}
