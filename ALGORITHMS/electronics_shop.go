// THIS IS A HACKERRANK TASK.
package main

func getMoneySpent(keyboards []int32, drives []int32, b int32) int32 {
    i:= 0
    most_expensive := int32(0)
    for i < len(keyboards){
        j := 0
        for j < len(drives){
            price := keyboards[i] + drives[j]
            if price <= b && price > most_expensive{
                most_expensive = price
            }
            j += 1    
        }
        i += 1
    }
    if most_expensive > 0{
        return most_expensive
    }
    return -1

}
