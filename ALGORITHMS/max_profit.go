// Initial version. Didn't pass all the tests.

package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}

func maxProfit(prices []int) int {
	buyIndex := 0
	sellIndex := 0
	buyVal := prices[0]
	sellVal := 0
	for i := range prices {
		if prices[i] > sellVal {
			sellIndex = i
			sellVal = prices[i]
			fmt.Println("Sell val", sellVal)
			fmt.Println("Sell index", sellIndex)
		}
		if prices[i] < buyVal {
			buyIndex = i
			buyVal = prices[i]
			fmt.Println("Buy Val", buyVal)
			fmt.Println("Buy index", buyIndex)
		}
		if buyIndex > sellIndex {
			sellVal = 0
			sellIndex = 0
		}

	}

	if sellVal != 0 {
		return (sellVal - buyVal)
	}
	return 0
}


//Second implementation. Still did not pass the test cases

func maxProfit2(prices []int) int {
	i := 0
	j := i + 1
	minVal := prices[0]
	maxVal := 0
	maxProfit := 0
	// buyIndex := 0
	// sellIndex := 0
	for i < j && j < len(prices) {
		if prices[j] > maxVal {
			//sellIndex = 0
			maxVal = prices[j]
		}

		if prices[i] < minVal {
			//buyIndex = i
			minVal = prices[i]
		}
		if maxVal-minVal > maxProfit {
			maxProfit = maxVal - minVal
		}
		// if buyIndex > sellIndex {
		// 	maxVal = 0
		// }

		i++
		j++
	}
	return maxProfit
}

//Third update. Finally passed, had to get assistance though. Seems like I wasn't getting the problem description correctly.
func maxProfit3(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	minPrice := prices[0]
	maxProfit := 0
	for i := range prices {
		currentPrice := prices[i]
		if currentPrice < minPrice {
			minPrice = currentPrice
		} else {
			currentProfit := currentPrice - minPrice
			if currentProfit > maxProfit {
				maxProfit = currentProfit
			}
		}

	}
	return maxProfit
}
