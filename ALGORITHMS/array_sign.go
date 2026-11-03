package main
// Second implementation. Logic was AI assisted. 0ms runtime and 4.94MB memory use.
func arraySign(nums []int) int {
    product := 1
    for _, num := range nums {
        if num == 0 {
            return 0
        }
    if num < 0 {
        product *= -1
    } else if product > 0 {
        continue
    }
    }
    return product
}


