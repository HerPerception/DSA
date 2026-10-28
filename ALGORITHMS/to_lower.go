package main

// First implementation. 0ms runtime and 4.18MB memory use.
func toLowerCase(s string) string {
    text := ""
    for _, ch := range s {
        if ch >= 'A' && ch <= 'Z' {
            ch += 32
        }
        text += string(ch)
    }
    return text
}
