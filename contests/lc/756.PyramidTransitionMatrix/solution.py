from collections import defaultdict
from typing import List


class Solution:
    def pyramidTransition(self, bottom: str, allowed: List[str]) -> bool:
        pairs = defaultdict(list)
        for rule in allowed:
            pairs[rule[:2]].append(rule[2])

        memo = {}

        def can_build(row: str) -> bool:
            if len(row) == 1:
                return True
            if row in memo:
                return memo[row]

            next_row = [""] * (len(row) - 1)

            def build_next(pos: int) -> bool:
                if pos == len(row) - 1:
                    return can_build("".join(next_row))

                pair = row[pos : pos + 2]
                if pair not in pairs:
                    return False
                for ch in pairs[pair]:
                    next_row[pos] = ch
                    if build_next(pos + 1):
                        return True
                return False

            memo[row] = build_next(0)
            return memo[row]

        return can_build(bottom)
