from typing import List


MOD = 1_000_000_007


class Solution:
    def countPermutations(self, complexity: List[int]) -> int:
        root = complexity[0]
        for value in complexity[1:]:
            if value <= root:
                return 0

        # Factorial of (n - 1) modulo MOD.
        fact = 1
        for x in range(2, len(complexity)):
            fact = (fact * x) % MOD
        return fact
