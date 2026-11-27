package main

// First implementation. 2ms runtime and 5.54MB memory use.
func canAliceWin(nums []int) bool {
    single := 0
    double := 0
    for _, num := range nums {
        if num < 10 {
            single += num
        } else {
            double += num
        }
    }
    if single > double || double > single {
        return true
    }
    return false
}
