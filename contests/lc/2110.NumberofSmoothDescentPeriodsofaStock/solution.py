from typing import List
from itertools import groupby
from math import comb
from operator import sub
from itertools import pairwise, starmap

class Solution:
    def getDescentPeriods(self, prices: List[int]) -> int:
        total = 0
        run_length = 0

        for i, price in enumerate(prices):
            if i > 0 and prices[i - 1] - price == 1:
                run_length += 1
            else:
                # Close the previous run and start a new one.
                total += run_length * (run_length + 1) // 2
                run_length = 1

        total += run_length * (run_length + 1) // 2
        return total

class Solution2:
    def getDescentPeriods(self, a: List[int]) -> int:
        return sum(comb(sum(g)+1,2) for d,g in groupby(map(sub,a,a[1:])) if d==1)+len(a)
        return sum(comb(sum(1 for _ in g)+1,2) for d,g in groupby(starmap(sub,pairwise(a))) if d==1)+len(a)