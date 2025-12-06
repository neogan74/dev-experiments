from collections import deque
from typing import List


class Solution:
    def countPartitions(self, nums: List[int], k: int) -> int:
        MOD = 1_000_000_007
        n = len(nums)
        dp = [0] * (n + 1)  # dp[i]: ways to partition first i elements
        pref = [0] * (n + 1)  # pref[i]: sum(dp[0..i])
        dp[0] = 1
        pref[0] = 1

        max_q: deque[int] = deque()
        min_q: deque[int] = deque()
        left = 0

        for i, val in enumerate(nums):
            while max_q and nums[max_q[-1]] < val:
                max_q.pop()
            max_q.append(i)

            while min_q and nums[min_q[-1]] > val:
                min_q.pop()
            min_q.append(i)

            # Shrink window until it satisfies max - min <= k
            while nums[max_q[0]] - nums[min_q[0]] > k:
                if max_q[0] == left:
                    max_q.popleft()
                if min_q[0] == left:
                    min_q.popleft()
                left += 1

            # Valid segment starts are in [left, i]; translate to dp indices
            total = pref[i]
            if left > 0:
                total -= pref[left - 1]
            dp[i + 1] = total % MOD
            pref[i + 1] = (pref[i] + dp[i + 1]) % MOD

        return dp[n] % MOD
