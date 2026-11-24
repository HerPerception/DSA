package main

/ *==================================================
THIS IS A HACKERRANK TASK.
*====================================================
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER_ARRAY ar
 */

func sockMerchant(n int32, ar []int32) int32 {
    sockMap := make(map[int32]int32)
    for _, sock := range ar {
        sockMap[sock]++
    }
    count := int32(0)
    for _, pairs := range sockMap {
        if pairs%2 == 0{
            num := pairs / 2
            count += num
        } else if pairs%2 != 0 && pairs > 0 {
            num := int(pairs/2)
            count += int32(num)
        }
    }
    return count
}
