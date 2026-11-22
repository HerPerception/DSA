package main

// First implementation. Suggested by AI, implementation by Ai. 0ms runtime and 5.70MB memory use.
func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
    count := 0
    for _, num := range arr1 {
        far := true
        for j := 0; j < len(arr2); j++ {
            dis := max(num, arr2[j]) - min(num, arr2[j])
            if dis <= d {
                far = false
                break
            }
        }
        if far {
        count++
        }
    }
    return count
}
