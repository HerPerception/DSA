package main

//First implementation. My own logic. 0ms runtime and 4.47MB memory use.
func threeConsecutiveOdds(arr []int) bool {
    count := 0
    for _, num := range arr {
        if num%2 != 0 {
            count++
            if count == 3 {
                return true
            }
        } else {
            count = 0
        }
    }
    return false
}
