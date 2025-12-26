class Solution:
    def min_deletion_size(self, strs):
        if not strs:
            return 0

        rows = len(strs)
        cols = len(strs[0])
        ans = 0

        for j in range(cols):
            for i in range(1, rows):
                if strs[i][j] < strs[i-1][j]:
                    ans += 1
                    break
        return ans