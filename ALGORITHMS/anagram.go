package main
//First implementation. My own logic. 7ms runtime..., 6.58MB memory use.
func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    anagramMapS := make(map[rune]int)
    anagramMapT := make(map[rune]int)
    anagramSlice := []rune{}
    for _, ch := range s {
        anagramMapS[ch]++
        anagramSlice = append(anagramSlice, ch)

    }
    for _, ch := range t {
        anagramMapT[ch]++
    }
    for _, eachVal := range anagramSlice {
        if anagramMapS[eachVal] != anagramMapT[eachVal] {
            return false
        }
    }
    return true
}

//Second implementation. Got steps from AI, not code. 5ms runtime..., 4.76MB memory use.
func isAnagram2(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    anagramMap := make(map[rune]int)
    for _, ch := range s {
        anagramMap[ch]++
    }
    for _, ch := range t {
        anagramMap[ch]--
        if anagramMap[ch] < 0 {
            return false
        }
    }
    return true
}
//Third implementation. Got steps from AI, not code. 0ms runtime and 4.75MB memory use.
func isAnagram3(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    anagramSlice := make([]int, 26)
    n := len(s)
    for i := 0; i < n; i++ {
        val := int(rune(s[i]) - 'a')
        anagramSlice[val] += 1
        val = int(rune(t[i]) - 'a')
        anagramSlice[val] -= 1
    }

    for _, num := range anagramSlice {
        if num != 0 {
            return false
        }
    }
    return true
}
