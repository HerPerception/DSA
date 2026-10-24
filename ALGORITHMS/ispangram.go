package main

//First implementation. My very own logic. 0ms runtime and 4.10MB
func checkIfPangram(sentence string) bool {
    alphabetArr := make([]int, 26)
    for _, ch := range sentence {
        if ch >= 'A' && ch <= 'Z' {
            in := int(ch - 'A')
            alphabetArr[in] += 1
        } else if ch >= 'a' && ch <= 'z' {
            in := int(ch - 'a')
            alphabetArr[in] += 1
        }
    }

    for _, num := range alphabetArr {
        if num <= 0 {
            return false
        }
    }
    return true
}
