# THIS IS A HACKERRANK TASK.
def pickingNumbers(a):
    # Step 1: Sort the array from smallest to largest
    a.sort()
    
    # Step 2: Create tracking variables
    start = 0
    max_length = 0
    
    # Step 3: Loop through the array with the 'end' pointer
    for end in range(len(a)):
        # Shrink the group if the difference is greater than 1
        while a[end] - a[start] > 1:
            start += 1
            
        # Calculate current group size and update record
        current_length = end - start + 1
        if current_length > max_length:
            max_length = current_length
            
    # Step 4: Return the maximum length found
    return max_length
