# THIS IS A HACKERRANK TASK.

# Completed the 'migratoryBirds' function below.
#
# The function is expected to return an INTEGER.
# The function accepts INTEGER_ARRAY arr as parameter.

def migratoryBirds(arr):
    id_dict = {}
    for val in arr:
        count = id_dict.get(val, 0)
        id_dict[val] = count + 1
    max_val = 0
    max_list = []
    for key, each in id_dict.items():
        if each > max_val:
            max_list = []
            max_val = each
            max_list.append(key)
        elif each == max_val:
            max_list.append(key)
            
    return min(max_list)
