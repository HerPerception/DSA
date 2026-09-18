package main

func strStr(haystack string, needle string) int {
    if haystack == needle {
        return 0
    }
    length := len(haystack)-len(needle)
	for in := 0; in <= length; in++ {
       n := in+len(needle)
		if haystack[in] == needle[0] && haystack[in:n] == needle{
			return in
		} 
	}
	return -1
}
