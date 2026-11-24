# THIS IS A HACKERRANK TASK.
def sockMerchant(n, ar):
    sockMap = {}
    for sock in ar:
        sockMap[sock] = sockMap.get(sock, 0) + 1

    count = 0
    for sock_count in sockMap.values():
        count += sock_count // 2
        
    return count
