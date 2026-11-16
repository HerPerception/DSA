# First implementation. My own logic. Debugging assisted by AI. 1ms runtime and 19.28MB memory use.
class Solution:
    def digitSum(self, s: str, k: int) -> str:
       
        while len(s) > k:
            s_list = []
            result = []
            for i in range(0, len(s), k):
                s_list.append(s[i:i+k])
            for idx in range(len(s_list)):
                total = 0
                for val in s_list[idx]:
                      total += int(val)
                s_list[idx] = str(total)
                result.append(str(total))
            s = "".join(result)
        return s
