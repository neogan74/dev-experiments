class Solution:
    def maxOperations(self, s: str) -> int:
        ones = [len(i) for i in s.split('0') if i]
        if not ones:
            return 0
        tot = 0
        ans = 0
        for f in ones[:-1]:
            tot += f
            ans += tot
        if s[-1] == '1':
            return ans
        return ans + tot + ones[-1]