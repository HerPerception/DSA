# THIS TASK IS ON HACKERRANK. My logic. Debugging assisted by AI.
if __name__ == '__main__':
    n = int(input())
    arr = map(int, input().split())
    arr = list(arr)
    arr.sort()
    first = arr[len(arr)-1]
    for i in range(len(arr)-1, -1, -1):
        if arr[i] < first:
            print(arr[i])
            break

# Second implementation. Logic by AI
def runner_up():
    n = int(input())
    arr = map(int, input().split())
    arr = set(arr)
    arr = list(arr)
    arr.sort()
    print(arr[-2])
   
