class Solution:
        
    def maxRunTime(self, n: int, batteries: List[int]) -> int:
        m = len(batteries)
        if n==m:
            return min(batteries)
        if m < n:
            return 0
        batteries.sort()
        S = sum(batteries[:-n]) # extra battery life to distribute to maximize time 
        L = batteries[-n:] 
        for i in range(n-1):
            if S < (i+1) * (L[i+1] - L[i]):
                return L[i] + S // (i+1)
            S -= (i+1) * (L[i+1] - L[i])
        return L[-1] + S // n