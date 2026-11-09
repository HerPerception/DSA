package main

// THIS SOLVES THE HACKERRANK TASK: APPLE AND ORANGE. MY OWN LOGIC.
// s and t is the range of Sam's house. The apple tree is located at a while the orange tree is located at b.
// The slices of ints apples and oranges contain the distances the apple and orange fruits fall from the tree. 
// The task is to find how many apple and orange fruits land on Sam's house.
func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
    countApple := 0
    countOrange := 0
    for dist := range apples {
        apples[dist] += a
        if apples[dist] >= s && apples[dist] <= t {
            countApple++
        }
    }
        for dist := range oranges {
        oranges[dist] += b
        if oranges[dist] >= s && oranges[dist] <= t {
            countOrange++
        }
        }
        fmt.Println(countApple)
        fmt.Println(countOrange)
}
