package main

package main

import (
    "fmt"
)

/*
 * Complete the 'dayOfProgrammer' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts INTEGER year as parameter.
 */

// THIS IS A HACKERRANK TASK.
// First implementation. Did not pass all the test cases.
func dayOfProgrammer(year int32) string {
    if year%4 == 0 {
        return fmt.Sprintf("12.09.%v", year)
    }
        return fmt.Sprintf("13.09.%v", year)
}
