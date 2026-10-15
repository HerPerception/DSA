package main

//First implementation. AI assisted.
func spiralOrder(matrix [][]int) []int {
    if len(matrix) == 0 { return []int{} }

    top, bottom := 0, len(matrix)-1
    left, right := 0, len(matrix[0])-1
    res := []int{}

    for top <= bottom && left <= right {
        // right
        for i := left; i <= right; i++ {
            res = append(res, matrix[top][i])
        }
        top++

        // down
        for i := top; i <= bottom; i++ {
            res = append(res, matrix[i][right])
        }
        right--

        // left
        if top <= bottom {
            for i := right; i >= left; i-- {
                res = append(res, matrix[bottom][i])
            }
            bottom--
        }

        // up
        if left <= right {
            for i := bottom; i >= top; i-- {
                res = append(res, matrix[i][left])
            }
            left++
        }
    }
    return res
}
