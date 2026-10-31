package main

// First implementation. My very own logic, debugging assisted by AI. 0ms runtime and 3.71MB memory use.
func subtractProductAndSum(n int) int {
    sum_total := 0
    product := 1
    for n > 0 {
        re := n%10
        sum_total += re
        product *= re
        n /= 10
    }
    return product - sum_total
}
