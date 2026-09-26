package main

func gradingStudents(grades []int32) []int32 {
    // Write your code here
    for i := range grades {
        if grades[i] < 38 {
            continue
        }
        multiple := ((grades[i] + 4) / 5) * 5
        if multiple - grades[i] < 3 {
            grades[i] = multiple
        }
    }
    return grades
}
