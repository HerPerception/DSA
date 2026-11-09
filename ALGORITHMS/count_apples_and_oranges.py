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
