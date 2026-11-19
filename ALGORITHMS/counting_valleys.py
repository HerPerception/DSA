# THIS IS A HACKERRANK TASK.

# The function is expected to return an INTEGER.
# The function accepts following parameters:
#  1. INTEGER steps
#  2. STRING path

def countingValleys(steps, path):
    level = 0
    valley = 0
    for step in path:
        if step =='D':
            level -= 1
        elif step =='U':
            level += 1
            if level == 0:
                valley += 1

    return valley
