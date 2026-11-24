package main

// First implementation. 0ms runtime and 5.02MB memory use.
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
}


func maxLength(nums []int) int {
    max_len := 0
    for i := 0; i < len(nums); i++ {
        prod := nums[i]
        g := nums[i]
        l := nums[i]
        for j :=i; j < len(nums); j++ {
            if prod > 25200 {
                break
            }
            if j > i {
                prod *= nums[j]
                g = gcd(g, nums[j])
                l = lcm(l, nums[j])
            }
            if prod == l*g {
                max_len = max((j-i+1), max_len)
            }
        }
    }
    return max_len
}
