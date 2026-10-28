package main

//My first implementation. My own logic. 1ms runtime and 6.71MB memory use.
func differenceOfSum(nums []int) int {
    total := 0
    digitSum := 0
    for _, num := range nums {
        total += num
        for num > 0 {
            n := num % 10
            digitSum += n
            num /= 10
        }
    }
    absVal := total - digitSum
    if absVal < 0 {
        absVal *= -1 
    }
    return absVal
}
