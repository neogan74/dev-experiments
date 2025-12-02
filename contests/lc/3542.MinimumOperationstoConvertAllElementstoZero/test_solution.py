import random
from functools import lru_cache
from unittest import TestCase

from solution import Solution


def brute_min_operations(nums):
    n = len(nums)

    @lru_cache(maxsize=None)
    def dfs(state):
        if all(value == 0 for value in state):
            return 0

        best = float("inf")
        for i in range(n):
            for j in range(i, n):
                sub = state[i : j + 1]
                sub_min = min(sub)
                if sub_min == 0:
                    continue  # operation does nothing if zeros already present

                new_state = list(state)
                for k in range(i, j + 1):
                    if state[k] == sub_min:
                        new_state[k] = 0

                best = min(best, 1 + dfs(tuple(new_state)))

        return best

    return dfs(tuple(nums))


class TestSolution(TestCase):
    def setUp(self):
        self.solution = Solution()

    def test_examples(self):
        examples = [
            ([0, 2], 1),
            ([3, 1, 2, 1], 3),
            ([1, 2, 1, 2, 1, 2], 4),
        ]
        for nums, expected in examples:
            with self.subTest(nums=nums):
                self.assertEqual(expected, self.solution.minOperations(nums[:]))

    def test_all_zero_array(self):
        nums = [0, 0, 0, 0]
        self.assertEqual(0, self.solution.minOperations(nums))

    def test_random_small_arrays_against_bruteforce(self):
        random.seed(0)
        for _ in range(200):
            n = random.randint(1, 5)
            nums = [random.randint(0, 3) for _ in range(n)]
            expected = brute_min_operations(nums)
            self.assertEqual(expected, self.solution.minOperations(nums[:]))
