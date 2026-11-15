package main

// AI assisted. 2ms runtime and 12.52MB memory use.
func numOfSubarrays(arr []int) int {
    runningSum := 0
    oddCount := 0
    evenCount := 1
    totalSubArrays := 0
    const modulo int = 1000000007
    for _, num := range arr {
        runningSum += num
        if runningSum%2 == 0 {
            totalSubArrays += oddCount
            evenCount += 1
        }else {
            totalSubArrays += evenCount
            oddCount += 1
        }
        totalSubArrays %= modulo

    }
    return totalSubArrays
}
