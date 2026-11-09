package main

// To reverse an array or slice
func reverseArray(a []int32) []int32 {
    revA := []int32{}
    for i := len(a)-1; i >= 0; i-- {
        revA = append(revA, a[i])
    }
    return revA
}
