# THIS IS A HACKERRANK TASK.
# Completed the 'dayOfProgrammer' function below.
#
# The function is expected to return a STRING.
# The function accepts INTEGER year as parameter.

#First implementation. Did not pass all the test cases.
def dayOfProgrammer(year):
    if year%4 == 0:
        return (f"12.09.{year}")
    else:
        return (f"13.09.{year}")


# Second implementation. AI assisted. Passed all test cases.
def dayOfProgrammer(year: int) -> str:
    if year == 1918:
        return "26.09.1918"
    elif year < 1918 and year % 4 == 0:
        return f"12.09.{year}"
    elif year > 1918 and (year % 400 == 0 or (year % 4 == 0 and year % 100 != 0)):
        return f"12.09.{year}"
    return f"13.09.{year}"
