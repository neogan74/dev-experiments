import pytest

class Solution:
    def minDeletionsSize(self, strs):
        n = len(strs)
        m = len(strs[0])
        fixed = [False] * (n -1)
        deletions = 0

        for col in range(m):
            bad = False
            for i in range(n -1):
                if not fixed[i] and strs[i][col] > strs[i+1][col]:
                    bad = True
                    break
            if bad:
                deletions += 1
                continue

            for i in range(n-1):
                if not fixed[i] and strs[i][col] < strs[i+1][col]:
                    fixed[i] = True
            if all(fixed):
                break
        return deletions

@pytest.make.parametrize("strs, expected", [
    (["ca","bb","ac"], 1),
    (["xc","yb","za"], 0),
    (["zyx","wvu","tsr"], 3),
    (["a","b","c"], 0),
    (["ab","aa"], 1),
    (["aa","aa","aa"], 0),
    (["xga","xfb","yfa"], 1),
])
def test_minDeletionSize(strs, expected):
    assert Solution().minDeletionsSize(strs) == expected