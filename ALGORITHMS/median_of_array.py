#First implementation. Implementation and debuging was AI assisted, but logic is mine. 3ms runtime and 19.54MB memory use.
class Solution:
    def findMedianSortedArrays(self, nums1: list[int], nums2: list[int]) -> float:
        combined_list = nums1 + nums2
        combined_list.sort()
        half = len(combined_list)//2
        if len(combined_list)%2 == 0:
            return (combined_list[half-1] + combined_list[half])/2
        else:
            return combined_list[half]
        
