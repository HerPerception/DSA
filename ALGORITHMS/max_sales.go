package main

// AI. Yet to pass the test cases.
func maximumSaleItems(items [][]int, budget int) int {
	n := len(items)
	freq := make([]int, n+1)
	minPrice := items[0][1]

	// 1. Count factor frequencies and find the global minimum price
	for _, item := range items {
		freq[item[0]]++
		if item[1] < minPrice {
			minPrice = item[1]
		}
	}

	// 2. Build (price, count of 2-item blocks) pairs for each item
	type bundle struct {
		price int
		count int
	}
	bundles := make([]bundle, 0, n)

	for _, item := range items {
		factor, price := item[0], item[1]
		multiples := 0
		for m := factor; m <= n; m += factor {
			multiples += freq[m]
		}
		// Exclude the item itself to get the number of valid freebie targets
		if multiples > 1 {
			bundles = append(bundles, bundle{price: price, count: multiples - 1})
		}
	}

	// 3. Sort 2-item blocks from cheapest to most expensive
	slices.SortFunc(bundles, func(a, b bundle) int {
		return a.price - b.price
	})

	// 4. Greedily buy 2-item blocks and check total items at each step
	remBudget := budget
	itemsSoFar := 0
	best := budget / minPrice // Baseline: spend everything on 1-item blocks

	for _, b := range bundles {
		// How many 2-item blocks from this bundle can we afford?
		canBuy := min(b.count, remBudget/b.price)

		remBudget -= canBuy * b.price
		itemsSoFar += canBuy * 2

		// Check total if we spend all remaining budget on 1-item blocks
		best = max(best, itemsSoFar+remBudget/minPrice)

		// If we couldn't buy the entire bundle, we don't have enough money for pricier ones
		if canBuy < b.count {
			break
		}
	}

	return best
}
