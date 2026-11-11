package main

// THIS IS A HACKERRANK TASK.
//  My logic.
func divisibleSumPairs(n int32, k int32, ar []int32) int32 {
    count := int32(0)
    for i :=int32(0); i < n-1; i++ {
        for j := i+1; j < n; j++{
            if (ar[i] + ar[j])%k == 0{
                count += 1
            }
        }
    }
    return count 

}
