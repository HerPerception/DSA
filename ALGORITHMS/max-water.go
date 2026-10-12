package main

// FIRST IMPLEMENTATION. 2ms RUNTIME
func findMin(x, y int)int{
    if x < y {
        return x
    }
    return y
}

func maxArea(height []int) int {
    left := 0
    right := len(height)-1
    maxWater := 0
    for left < right {
        minNum := findMin(height[left], height[right])
        area := (right - left) * minNum
        if area > maxWater {
            maxWater = area
        }
        if height[left] < height[right] {
            left++
        } else {
            right--
        } 
    }
    return maxWater
}
