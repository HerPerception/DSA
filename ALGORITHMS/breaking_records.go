package main

// THIS IS A HACKERRANK TASK.
// AI assisted.
func breakingRecords(scores []int32) []int32 {
    min_score := scores[0]
    max_score := scores[0]
    min_breaks := int32(0)
    max_breaks := int32(0)
    for i := 1; i < len(scores); i++ {
        if scores[i] > max_score {
            max_score = scores[i]
            max_breaks += 1
        } else if scores[i] < min_score {
            min_score = scores[i]
            min_breaks += 1
        }
    }
    return append([]int32{}, max_breaks, min_breaks)
}
