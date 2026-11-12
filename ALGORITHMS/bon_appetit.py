# THIS IS A HACKERRANK TASK.
# Logic is mine.
def bonAppetit(bill, k, b):
    total = sum(bill)
    anna = (total - bill[k])/2
    if anna == b:
        print("Bon Appetit")
    else:
        print(int(b - anna))
