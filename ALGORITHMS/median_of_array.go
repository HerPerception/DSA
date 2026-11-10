package main

// First implementation. 3ms runtime and 7.21MB memory use.
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    combined_list := append(nums1, nums2...)
    sort.Ints(combined_list)
    half := int(len(combined_list)/2)
    if len(combined_list)%2 == 0{
        return float64((combined_list[half-1] + combined_list[half]))/2
    }
    return float64(combined_list[half])
        
}
