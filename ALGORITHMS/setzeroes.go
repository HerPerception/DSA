package main

//First implementation. 44ms runtime and 10MB memory use.
func setZeroes(matrix [][]int)  {
   indexSlice := [][]int{}
   for i, nums := range matrix {
    for j, num := range nums {
        if num == 0 {
            indexSlice = append(indexSlice, []int{i, j})
        }
    }
   } 

   for _, index := range indexSlice {
     i := index[0]
     j := index[1]
     for in, val := range matrix {
        if in == i {
            for idx := range val {
                val[idx] = 0
            }
        }
        val[j] = 0 
     }
   }
 fmt.Println(matrix)
}
