package main
import "strings"
//**First implementation. Yet to pass. 18-10-2026, 11:57PM**
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

//Eventually corrected the  first implementation to pass all test cases but with 108ms runtime and 17.47MB memory use.
//19-10-2026, 12:01AM.
func reverseWordsCorrected(str string) string {
    s := strings.Fields(str)
    reversedStr :=  ""
    for i, eachword := range s {
        text := ""
        for i := len(eachword)-1; i >= 0; i-- {
            text += string(eachword[i])
        }
        reversedStr += text
        if i != len(s)-1 {
         reversedStr += " "
        }
    }
    return reversedStr
}

//2nd implementation. Passed all test cases. 3ms runtime and 8.09MB memory use.
func reverseWords2(s string) string {
  str := strings.Fields(s)
  var revStr strings.Builder
  for in, eachWord := range str {
    var revWord strings.Builder
    for i := len(eachWord)-1; i >= 0; i-- {
        revWord.WriteByte(eachWord[i])
    }
    revStr.WriteString(revWord.String())
    if in != len(str)-1 {
        revStr.WriteString(" ")
    }
  }
  return revStr.String()
}
