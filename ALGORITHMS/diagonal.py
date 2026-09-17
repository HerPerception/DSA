#!/bin/python3

import math
import os
import random
import re
import sys

#
# Complete the 'diagonalDifference' function below.
#
# The function is expected to return an INTEGER.
# The function accepts 2D_INTEGER_ARRAY arr as parameter.
#

def diagonalDifference(arr):
    # Write your code here
    left = 0
    right = 0
    idx = 0
    for each_arr in arr:
        index = len(each_arr)-1
        print(each_arr[idx])
        left += each_arr[idx]
        idx += 1
    for each_arr in arr:
        right += each_arr[index]
        print(each_arr[index])
        index -= 1

    num = left - right
    if num < 0:
        num *= -1
        
    return num
        
        
        

if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    n = int(input().strip())

    arr = []

    for _ in range(n):
        arr.append(list(map(int, input().rstrip().split())))

    result = diagonalDifference(arr)

    fptr.write(str(result) + '\n')

    fptr.close()
