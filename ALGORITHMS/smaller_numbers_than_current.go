package main

//First implementation. My logic. 4ms runtime and 4.98MB memory use.
func smallerNumbersThanCurrent(nums []int) []int {
    countArr := []int{}
    for _, num := range nums {
            count := 0
        for _, n := range nums {
            if num > n {
                count += 1
            }
        }
        countArr = append(countArr, count)
    }
    return countArr
}
