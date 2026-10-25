package main

//First implementation. My very own logic. 0ms runtime and 5.51MB memory use.
func mostWordsFound(sentences []string) int {
    maxWords := 0
    for _, sentence := range sentences {
        count := 0
        for _, ch := range sentence {  
            if ch == ' ' {
                count += 1
            }
        }
        count += 1
        if count > maxWords {
            maxWords = count
        }
    }
    return maxWords
}
