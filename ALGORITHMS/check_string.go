package main

// First implementation. Logic is mine. 0ms runtime and 4.66MB memory use.
func checkString(s string) bool {
    b_seen := false
    for _, ch := range s {
        if ch == 'b' {
            b_seen = true
        }
        if ch == 'a' && b_seen == true {
            return false
        }
    }
    return true
}

// Second implementation. Suggested by AI. 0ms runtime and 4.69MB memory use.
func checkString2(s string) bool {
    for i := 0; i < len(s)-1; i++ {
        if s[i] == 'b' && s[i+1] == 'a'{
            return false
        }
    }
        return true
}
