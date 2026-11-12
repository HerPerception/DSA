package main

// THIS IS A HACKERRANK TASK.
func bonAppetit(bill []int32, k int32, b int32) {
    total := int32(0)
    for _, val := range bill {
        total += val
    }
    anna := (total - bill[k])/2
    if anna == b{
        fmt.Println("Bon Appetit")
    }else{
        fmt.Println(b - anna)
    }
}
