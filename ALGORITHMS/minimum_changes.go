package main

// First implementation. 91ms runtime and 9.90MB memory use. Logic and implementation assisted by AI.
func costHelper(substr string) int {
    d := []int{}
    bestScore := 999999
    n := len(substr)
    for i := 1; i < n; i++ {
        if n%i == 0{
            d = append(d, i)
        }
    }
    for _, step := range d {
        fixCosts := 0
        for start := 0; start < step; start++ {
        left, right := start, n - step + start
        for left < right {
            if substr[left] != substr[right]{
                fixCosts += 1
            }
            left += step
            right -= step
            }
        }
        if fixCosts < bestScore {
            bestScore = fixCosts
        }
    }     
    return bestScore
}

func minimumChanges(s string, k int) int {
	n := len(s)

	// 1. Pre-calculate the fix cost for every valid substring (length >= 2)
	cost := make([][]int, n)
	for i := 0; i < n; i++ {
		cost[i] = make([]int, n)
		for j := i + 1; j < n; j++ {
			cost[i][j] = costHelper(s[i : j+1])
		}
	}

	// 2. Initialize DP table with a large number
	// dp[pieces][end] = min changes to split s[0...end] into 'pieces' substrings
	dp := make([][]int, k+1)
	for p := 0; p <= k; p++ {
		dp[p] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[p][j] = 999999
		}
	}

	// Base case: 1 piece from index 0 to j (must have at least 2 letters, so j >= 1)
	for j := 1; j < n; j++ {
		dp[1][j] = cost[0][j]
	}

	// 3. Build up from 2 pieces to k pieces
	for pieces := 2; pieces <= k; pieces++ {
		for end := 1; end < n; end++ {
			// Try every valid split point 'mid'
			// Leave at least 2 characters for the last piece (mid + 1 to end)
			for mid := 0; mid <= end-2; mid++ {
				if dp[pieces-1][mid]+cost[mid+1][end] < dp[pieces][end] {
					dp[pieces][end] = dp[pieces-1][mid] + cost[mid+1][end]
				}
			}
		}
	}

	return dp[k][n-1]
}
