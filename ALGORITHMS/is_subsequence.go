package main
//First implementation. My logic. 0ms runtime and 3.93 memory use
func isSubsequence(s string, t string) bool {
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
