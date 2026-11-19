// THIS IS A HACKERRANK TASK
package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

/*
 * Complete the 'countingValleys' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER steps
 *  2. STRING path
 */

func countingValleys(steps int32, path string) int32 {
    level := 0
    valley := int32(0)
    for _, step := range path {
        if step =='D' {
            level -= 1
        } else if step =='U' {
            level += 1
             if level == 0 {
            valley += 1
        }
        }
       
    }
    return valley
}
