package main

// First implementation. AI assisted. 0ms runtime and 4.65MB memory use.
func minimumCost(c1 int, c2 int, cB int, n1 int, n2 int) int64 {
    cost1, cost2, costBoth, need1, need2 := int64(c1), int64(c2), int64(cB), int64(n1), int64(n2)
    totalA:= (need1 * cost1) + (need2 * cost2)
    overlap := min(need1, need2)
    rem1 := need1 - overlap
    rem2 := need2 - overlap
    rem1 *= cost1
    rem2 *= cost2
    overlap *= costBoth
    totalB := overlap + rem1 + rem2
    totalC := max(need1, need2) * costBoth
    if totalA <= totalB && totalA <= totalC {
        return totalA
    } else  if totalB <= totalA && totalB <= totalC {
        return totalB
    }
    return totalC
}
