package main

// THIS IS A HACKERRANK TASK
// Completed the catAndMouse function below.
func catAndMouse(x int32, y int32, z int32) string {
    catA := max(x, z) - min(x, z)
    catB := max(y, z) - min(y, z)
    if catA > catB {
        return "Cat B"
    } else if catB > catA {
        return "Cat A"
    }
    return "Mouse C"
}
