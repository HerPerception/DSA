package main

/*==============================================================
First implementation. 203ms runtime and 31.16MB memory use.
=============================================================*/
package main

func findSubtreeSizes(parent []int, s string) []int {
	n := len(parent)

	// Step 1: Build initial tree (boss -> slice of direct reports)
	children := make([][]int, n)
	for i := 1; i < n; i++ {
		children[parent[i]] = append(children[parent[i]], i)
	}

	// Copy parent slice into newParent
	newParent := make([]int, n)
	copy(newParent, parent)

	// Track the closest ancestor for each letter ('a' to 'z'), initialized to -1
	var lastSeen [26]int
	for i := range lastSeen {
		lastSeen[i] = -1
	}

	// Step 2: Walk down the tree to find new bosses
	var dfs func(node int)
	dfs = func(node int) {
		charIdx := s[node] - 'a' // Convert 'a'-'z' to index 0-25

		if lastSeen[charIdx] != -1 {
			newParent[node] = lastSeen[charIdx]
		}

		seen := lastSeen[charIdx]
		lastSeen[charIdx] = node

		for _, child := range children[node] {
			dfs(child)
		}

		// Backtrack: restore the previous ancestor
		lastSeen[charIdx] = seen
	}

	dfs(0)

	// Step 3: Build the updated tree and count subtree sizes
	newChildren := make([][]int, n)
	for i := 1; i < n; i++ {
		newChildren[newParent[i]] = append(newChildren[newParent[i]], i)
	}

	answer := make([]int, n)
	for i := range answer {
		answer[i] = 1
	}

	var countSizes func(node int)
	countSizes = func(node int) {
		for _, child := range newChildren[node] {
			countSizes(child)
			answer[node] += answer[child]
		}
	}

	countSizes(0)
	return answer
}
