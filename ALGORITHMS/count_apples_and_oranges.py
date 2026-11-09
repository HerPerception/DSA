# THIS SOLVES THE HACKERRANK TASK: APPLE AND ORANGE. MY OWN LOGIC. 
# Was originally implemented in Golang. Debugging was AI assisted.
# s and t is the range of Sam's house. The apple tree is located at a while the orange tree is located at b.
# The slices of ints apples and oranges contain the distances the apple and orange fruits fall from the tree. 
# The task is to find how many apple and orange fruits land on Sam's house.
def countApplesAndOranges(s, t, a, b, apples, oranges):
    countApple = 0
    countOrange = 0
    for dist in range(len(apples)):
        apples[dist] += a
        if apples[dist] >= s and apples[dist] <= t:
            countApple += 1

    for dist in range(len(oranges)):
        oranges[dist] += b
        if oranges[dist] >= s and oranges[dist] <= t:
            countOrange += 1

    print(countApple)
    print(countOrange)
