import unittest

from solution import Solution


class MinOperationsTests(unittest.TestCase):
    def test_examples(self) -> None:
        solver = Solution()
        self.assertEqual(solver.minOperations([3, 9, 7], 5), 4)
        self.assertEqual(solver.minOperations([4, 1, 3], 4), 0)
        self.assertEqual(solver.minOperations([3, 2], 6), 5)

    def test_already_divisible(self) -> None:
        solver = Solution()
        self.assertEqual(solver.minOperations([2, 2, 2], 2), 0)

    def test_k_is_one(self) -> None:
        solver = Solution()
        self.assertEqual(solver.minOperations([10, 20, 30], 1), 0)


if __name__ == "__main__":
    unittest.main()
