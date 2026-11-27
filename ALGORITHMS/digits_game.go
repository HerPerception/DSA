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

// Second implementation. 0ms runtime and 5.44MB memory use.
func canAliceWin(nums []int) bool {
    total := 0
    for _, num := range nums {
        if num < 10 {
            total -= num
        } else {
            total += num
        }
    }
  return total != 0
}
