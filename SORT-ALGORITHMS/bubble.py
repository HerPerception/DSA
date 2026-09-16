my_array = [9, 5, 6, 4, 6, 1, 3, 10, 7, 2]

n = len(my_array)
for i in range(n-1):
    swapped = False
    for j in range(n-i-1):
        if my_array[j] > my_array[j+1]:
            my_array[j], my_array[j+1] = my_array[j+1], my_array>
            swapped = True
    if not swapped:
        break
print(my_array)
