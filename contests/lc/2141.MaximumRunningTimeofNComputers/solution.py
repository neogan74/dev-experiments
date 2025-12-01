from typing import List


class Solution:
    def maxRunTime(self, n: int, batteries: List[int]) -> int:
        # Example: n=2, batteries=[3,3,3]
        # total=9, hi=9//2=4. Check mid=3 -> feasible (sum min=6 >= 6).
        # Check mid=4 -> feasible (sum min=8 >= 8). Answer=4.
        #
        # Flow:
        # 1) Binary search on time in [0, total/n].
        # 2) A time t is feasible if sum(min(battery, t)) >= n * t.
        # 3) Return the largest feasible t.
        total = sum(batteries)
        lo, hi = 0, total // n

        def can_run(time: int) -> bool:
            energy = 0
            for charge in batteries:
                if charge >= time:
                    energy += time
                else:
                    energy += charge
            return energy >= n * time

        while lo < hi:
            mid = (lo + hi + 1) // 2
            if can_run(mid):
                lo = mid
            else:
                hi = mid - 1
        return lo
