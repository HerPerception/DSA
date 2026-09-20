func isPalindrome(x int) bool {
    s := strconv.Itoa(x)
    var reverse func (s string)string
    reverse = func (s string)string {
        if len(s) == 1 {
            return s
        }
        return reverse(s[1:]) + string(s[0])
    }
    return s == reverse(s)
}
