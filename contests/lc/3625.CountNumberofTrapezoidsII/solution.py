from math import gcd
from collections import defaultdict
from typing import List

class Solution:
    def countTrapezoids(self, points: List[List[int]]) -> int:
        t = defaultdict(lambda: defaultdict(int))
        v = defaultdict(lambda: defaultdict(int))

        n = len(points)

        for i in range(n):
            x1, y1 = points[i]
            for j in range(i + 1, n):
                x2, y2 = points[j]
                dx = x2 - x1
                dy = y2 - y1

                if dx < 0 or (dx == 0 and dy < 0):
                    dx = -dx
                    dy = -dy

                g = gcd(dx, abs(dy))
                sx = dx // g
                sy = dy // g

                des = sx * y1 - sy * x1

                key1 = (sx << 12) | (sy + 2000)
                key2 = (dx << 12) | (dy + 2000)

                t[key1][des] += 1
                v[key2][des] += 1

        return self.count(t) - self.count(v) // 2

    def count(self, mp):
        ans = 0

        for inner in mp.values():
            total = sum(inner.values())
            remaining = total

            for val in inner.values():
                remaining -= val
                ans += val * remaining

        return ans
    

    def countTrapezoids2(self, points: List[List[int]]) -> int:
        n = len(points)
        inf = 10**9 + 7
        slope_to_intercept = defaultdict(list)
        mid_to_slope = defaultdict(list)
        ans = 0
        
        for i in range(n):
            x1, y1 = points[i]
            for j in range(i + 1, n):
                x2, y2 = points[j]
                dx = x1 - x2
                dy = y1 - y2
                
                if x2 == x1:
                    k = inf
                    b = x1
                else:
                    k = (y2 - y1) / (x2 - x1)
                    b = (y1 * dx - x1 * dy) / dx
                
                mid = (x1 + x2) * 10000 + (y1 + y2)
                slope_to_intercept[k].append(b)
                mid_to_slope[mid].append(k)

        for sti in slope_to_intercept.values():
            if len(sti) == 1:
                continue
            
            cnt = defaultdict(int)
            for b_val in sti:
                cnt[b_val] += 1
            
            total_sum = 0
            for count in cnt.values():
                ans += total_sum * count
                total_sum += count

        for mts in mid_to_slope.values():
            if len(mts) == 1:
                continue
            
            cnt = defaultdict(int)
            for k_val in mts:
                cnt[k_val] += 1
            
            total_sum = 0
            for count in cnt.values():
                ans -= total_sum * count
                total_sum += count
        
        return ans