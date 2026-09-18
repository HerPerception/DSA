package main

func removeElement(nums []int, val int) int {
    count := 0
   for i := 0; i < len(nums); i++ {
    if nums[i] == val {
        nums = append(nums[:i], nums[i+1:]...)
        i--
    } else {
        count++
    }
   }
   return count 
}
