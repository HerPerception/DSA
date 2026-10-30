package main

// First implementation. My logic. 2ms runtime and 6.62MB memory use.
func findWordsContaining(words []string, x byte) []int {
    indexArr := []int{}
    for in, word := range words {
        for i := 0; i < len(word); i++ {
            if word[i] == x {
                indexArr = append(indexArr, in)
                break
            }
        }
    }
    return indexArr
}
