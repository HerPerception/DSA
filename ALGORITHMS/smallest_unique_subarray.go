package main
/*=====================================================================
    First implementation. Exceeded time limit for some of the test cases.
 =====================================================================*/
func smallestUniqueSubarray(nums []int) int {
    for i := 1; i <= len(nums); i++ {
        subMap := make(map[string]int)
        for j := 0; j <= len(nums)-i; j++ {
            subarr := nums[j:j+i]
            substr := fmt.Sprint(subarr)
            subMap[substr]++
        }
        for _, count := range subMap {
            if count == 1 {
            return i
            }
        }
    }
    return len(nums)
}


/*=====================================================
    Second implementation. AI assisted.
    Yet to fully understand the helper function.
=======================================================*/
func smallestUniqueSubarray(nums []int) int {
	low, high := 1, len(nums)
	ans := len(nums)

	for low <= high {
		mid := low + (high-low)/2
		if hasUnique(nums, mid) {
			ans = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return ans
}

func hasUnique(nums []int, length int) bool {
	const b1, b2 uint64 = 1000003, 1000033
	var h1, h2 uint64
	var pow1, pow2 uint64 = 1, 1

	for i := 0; i < length; i++ {
		val := uint64(nums[i] + 1)
		h1 = h1*b1 + val
		h2 = h2*b2 + val
		if i > 0 {
			pow1 *= b1
			pow2 *= b2
		}
	}

	counts := make(map[[2]uint64]int, len(nums)-length+1)
	counts[[2]uint64{h1, h2}]++

	for i := length; i < len(nums); i++ {
		out := uint64(nums[i-length] + 1)
		in := uint64(nums[i] + 1)

		h1 = (h1-out*pow1)*b1 + in
		h2 = (h2-out*pow2)*b2 + in

		counts[[2]uint64{h1, h2}]++
	}

	for _, c := range counts {
		if c == 1 {
			return true
		}
	}
	return false
}
