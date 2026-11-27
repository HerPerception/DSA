package main

/*=======================================================================================
First implementation. 15ms runtime and 9.13MB memory use.
=========================================================================================*/
func maximumSaleItems(items [][]int, budget int) int {
    minPrice := 1501
    freq := make([]int, 1501)
    for i := 0; i < len(items); i++ {
        factor := items[i][0]
        freq[factor]++
        if items[i][1] < minPrice {
            minPrice = items[i][1]
        }
       
    }
    values := make([]int, len(items))
    for j := 0; j < len(items); j++ {
        factor := items[j][0]
        count := 0
        for m := factor; m <= 1500; m += factor {
            count += freq[m]
        }
        values[j] = count
    }
    dp := make([]int, budget+1)
    for k := 0; k < len(items); k++ {
        price := items[k][1]
        val := values[k]
        for c := budget; c >= price; c-- {
            if dp[c-price] + val > dp[c] {
                dp[c] = dp[c-price] + val
            }
        }
    }
    ans := 0
    for c := 0; c <= budget; c++ {
        total := dp[c] + (budget - c)/ minPrice
        if total > ans {
            ans = total
        }
    }
    return ans
}
