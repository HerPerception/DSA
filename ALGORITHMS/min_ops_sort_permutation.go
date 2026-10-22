package main

import (
  "strings"
  "strconv"
  )

//First implementation. Wrote part of the code. AI assisted. 938 0f 999 test cases passed.
type  MinSteps struct {
    nums []int
    steps int
}

func arrayToStr(nums []int)string {
    var str strings.Builder
    for i, num := range nums {
    str.WriteString(strconv.Itoa(num))
    if i != len(nums)-1 {
    str.WriteString(",")
    }
    }
    return str.String()
}

func minOperations(nums []int) int {
    target := make([]int, len(nums))
    for i := range target {
        target[i] = i
    }
    same := true
    for i := 0; i < len(nums); i++ {
        if nums[i] != target[i] {
            same = false
            break
        }
    }
    if same == true {
        return 0
    }
    numsCopy := make([]int, len(nums))
    copy(numsCopy, nums)
    start := MinSteps{nums: numsCopy, steps: 0}
    queue := []MinSteps{}
    queue = append(queue, start)
    visited := make(map[string]bool)
    str := arrayToStr(start.nums)
    visited[str] = true 

    for len(queue) > 0 {
    current := queue[0]
    queue = queue[1:]

    // Check whether the current array is sorted.
    sorted := true
    for i := range target {
        if current.nums[i] != target[i] {
            sorted = false
            break
        }
    }

    if sorted {
        return current.steps
    }

    // Operation 1: Rotate left by one position.
    rotated := make([]int, len(current.nums))
    for i := 0; i < len(current.nums)-1; i++ {
        rotated[i] = current.nums[i+1]
    }
    rotated[len(current.nums)-1] = current.nums[0]

    rotatedKey := arrayToStr(rotated)
    if !visited[rotatedKey] {
        visited[rotatedKey] = true
        queue = append(queue, MinSteps{
            nums:  rotated,
            steps: current.steps + 1,
        })
    }

    // Operation 2: Reverse the entire array.
    reversed := make([]int, len(current.nums))
    for i := range current.nums {
        reversed[i] = current.nums[len(current.nums)-1-i]
    }

    reversedKey := arrayToStr(reversed)
    if !visited[reversedKey] {
        visited[reversedKey] = true
        queue = append(queue, MinSteps{
            nums:  reversed,
            steps: current.steps + 1,
        })
    }
}

return -1
}

