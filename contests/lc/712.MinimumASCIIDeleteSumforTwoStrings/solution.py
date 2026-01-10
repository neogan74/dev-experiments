class Solution:
    def minimumDeleteSum(self, s1: str, s2: str) -> int:
        m, n = len(s1), len(s2)
        dp = [0] * (n + 1)

        for j in range(n - 1, -1, -1):
            dp[j] = dp[j + 1] + ord(s2[j])

        for i in range(m - 1, -1, -1):
            prev_diag = dp[n]
            dp[n] = dp[n] + ord(s1[i])
            for j in range(n - 1, -1, -1):
                temp = dp[j]
                if s1[i] == s2[j]:
                    dp[j] = prev_diag
                else:
                    delete_s1 = ord(s1[i]) + temp
                    delete_s2 = ord(s2[j]) + dp[j + 1]
                    dp[j] = delete_s1 if delete_s1 < delete_s2 else delete_s2
                prev_diag = temp

        return dp[0]
