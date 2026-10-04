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
