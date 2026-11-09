package main

//First implementation. Logic and debugging was AI assisted. 0ms runtime and 4.14MB memory use.
func triangleType(nums []int) string {
    a := nums[0]
    b := nums[1]
    c := nums[2]
    if a + b <= c || a + c <= b || b + c <= a {
        return "none"
    } 
    if a == b && b == c {
        return "equilateral"
    } else if a == b || b == c  || a == c{
        return "isosceles"
    }
    return "scalene"
}
