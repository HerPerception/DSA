package main

import (
	"strconv"
	"strings"
)

func digitSum(s string, k int) string {
	for len(s) > k {
		var s_list []string
		var result []string

		for i := 0; i < len(s); i += k {
			end := i + k
			if end > len(s) {
				end = len(s)
			}
			s_list = append(s_list, s[i:end])
		}

		for idx := 0; idx < len(s_list); idx++ {
			total := 0
			for _, val := range s_list[idx] {
				num, _ := strconv.Atoi(string(val))
				total += num
			}
			s_list[idx] = strconv.Itoa(total)
			result = append(result, strconv.Itoa(total))
		}

		s = strings.Join(result, "")
	}
	return s
}
