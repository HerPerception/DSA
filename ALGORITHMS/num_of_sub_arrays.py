# First implementation. Yet to understand it. 105ms runtime and 24.68MB memory use.
class Solution:
    def numOfSubarrays(self, arr: list[int]) -> int:
        running_sum = 0
        odd_count = 0
        even_count = 1
        total_sub_arrays = 0
        MODULO = 1000000007
        for i in range(len(arr)):
            running_sum += arr[i]
            if running_sum%2 == 0:
                total_sub_arrays += odd_count
                even_count += 1
            else:
                total_sub_arrays += even_count
                odd_count += 1

            total_sub_arrays %= MODULO

        return total_sub_arrays
