def gradingStudents(grades):
    for index, grade in enumerate(grades):
        if grade < 38:
            continue
    # Python uses // for integer division
        next_multiple = ((grade + 4) // 5) * 5 
        if next_multiple - grade < 3:
            grades[index] = next_multiple
    return grades
