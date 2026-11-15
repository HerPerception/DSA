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
