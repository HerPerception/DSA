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
