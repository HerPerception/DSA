package main
/*============================================
THIS IS A HACKERRANK TASK.
=============================================*/

func getTotalX(a []int32, b []int32) int32 {
	// 1. Find the maximum value in a
	start := a[0]
	for _, num := range a {
		if num > start {
			start = num
		}
	}

	// 2. Find the minimum value in b
	end := b[0]
	for _, num := range b {
		if num < end {
			end = num
		}
	}

	// 3. Early exit if no number can fit between them
	if start > end {
		return 0
	}

	var count int32 = 0

	// 4. Test every candidate x from start to end
	for x := start; x <= end; x++ {
		valid := true

		// Check if every element in a is a factor of x
		for _, num := range a {
			if x%num != 0 {
				valid = false
				break
			}
		}

		// Only check b if a passed
		if valid {
			for _, num := range b {
				if num%x != 0 {
					valid = false
					break
				}
			}
		}

		if valid {
			count++
		}
	}

	return count
}
