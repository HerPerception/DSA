package main

/* THIS IS A HACKERRANK TASK.
kangaroo returns "YES" if two kangaroos starting at x1 and x2 with jump
rates v1 and v2 will ever land on the same location after the same number
of jumps, and "NO" otherwise.
*/
func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
    if v1 <= v2 {
        return "NO"
    }
    j := (x2 - x1) / (v1 - v2)
    if x1 + (j * v1) == x2 + (j * v2) {
        return "YES"
    }
    return "NO"
}
