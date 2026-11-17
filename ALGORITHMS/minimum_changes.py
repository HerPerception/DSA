#==============================================================================
# First implementation assisted by AI. 2698ms runtime and 19.65MB memory use.
#==============================================================================
class Solution:
    def costHelper(self, substr: str) -> int:
        n = len(substr)
        best_score = 999999
        
        # Find all valid divisors d (from 1 up to n - 1)
        divisors = [i for i in range(1, n) if n % i == 0]
        
        for step in divisors:
            fix_costs = 0
            # Check each repeating group starting at index 'start'
            for start in range(step):
                left = start
                right = n - step + start
                
                while left < right:
                    if substr[left] != substr[right]:
                        fix_costs += 1
                    left += step
                    right -= step
                    
            if fix_costs < best_score:
                best_score = fix_costs
                
        return best_score

    def minimumChanges(self, s: str, k: int) -> int:
        n = len(s)
        
        # 1. Pre-calculate the fix cost for every valid substring (length >= 2)
        cost = [[0] * n for _ in range(n)]
        for i in range(n):
            for j in range(i + 1, n):
                cost[i][j] = self.costHelper(s[i : j + 1])
                
        # 2. Initialize DP table with a large number
        # dp[pieces][end] = min changes to split s[0...end] into 'pieces' substrings
        dp = [[999999] * n for _ in range(k + 1)]
        
        # Base case: 1 piece from index 0 to j (must have at least 2 letters, so j >= 1)
        for j in range(1, n):
            dp[1][j] = cost[0][j]
            
        # 3. Build up from 2 pieces to k pieces
        for pieces in range(2, k + 1):
            for end in range(1, n):
                # Try every valid split point 'mid'
                # Leave at least 2 characters for the last piece (mid + 1 to end)
                for mid in range(end - 1):
                    if dp[pieces - 1][mid] + cost[mid + 1][end] < dp[pieces][end]:
                        dp[pieces][end] = dp[pieces - 1][mid] + cost[mid + 1][end]
                        
        return dp[k][n - 1]


#=====================================================================================================
# Second implementation, more efficient and provided by AI. 2167ms runtime and 19.66MB memory use. 
#=====================================================================================================
class Solution:
    def minimumChanges(self, s: str, k: int) -> int:
        n = len(s)

        # 1. Precompute divisors for every possible substring length (from 2 to n)
        divisors = [[] for _ in range(n + 1)]
        for length in range(2, n + 1):
            for d in range(1, length):
                if length % d == 0:
                    divisors[length].append(d)

        # 2. Helper to find min changes for s[start...end] without slicing strings
        def cost_helper(start: int, end: int) -> int:
            length = end - start + 1
            best_score = 999999

            for step in divisors[length]:
                fix_costs = 0
                for offset in range(step):
                    left = start + offset
                    right = end - step + 1 + offset

                    while left < right:
                        if s[left] != s[right]:
                            fix_costs += 1
                        left += step
                        right -= step

                if fix_costs < best_score:
                    best_score = fix_costs

            return best_score

        # 3. Pre-calculate the fix cost for every valid substring (length >= 2)
        cost = [[0] * n for _ in range(n)]
        for i in range(n):
            for j in range(i + 1, n):
                cost[i][j] = cost_helper(i, j)

        # 4. Initialize DP table with a large number
        # dp[pieces][end] = min changes to split s[0...end] into 'pieces' substrings
        dp = [[999999] * n for _ in range(k + 1)]

        # Base case: 1 piece from index 0 to j
        for j in range(1, n):
            dp[1][j] = cost[0][j]

        # 5. Build up from 2 pieces to k pieces
        for pieces in range(2, k + 1):
            # Minimum end index needed so every piece has at least 2 letters
            min_end = pieces * 2 - 1
            for end in range(min_end, n):
                # Minimum split index needed for the previous (pieces - 1) parts
                min_mid = (pieces - 1) * 2 - 1
                for mid in range(min_mid, end - 1):
                    total = dp[pieces - 1][mid] + cost[mid + 1][end]
                    if total < dp[pieces][end]:
                        dp[pieces][end] = total

        return dp[k][n - 1]
