package main
//First implementation. Yet to pass.
func reverseWords(str string) string {
    s := strings.Fields(str)
    reversedStr :=  ""
    for _, eachword := range s {
        if eachword == " " {
            reversedStr += " "
            continue
        }
        text := ""
        for i := len(eachword)-1; i >= 0; i-- {
            text += string(eachword[i])
        }
        reversedStr += text
    }
    return reversedStr
}
