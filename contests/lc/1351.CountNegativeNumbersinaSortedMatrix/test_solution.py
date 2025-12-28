import unittest

from solution import Solution


class TestCountNegatives(unittest.TestCase):
    def setUp(self) -> None:
        self.solution = Solution()

    def test_example1(self) -> None:
        grid = [
            [4, 3, 2, -1],
            [3, 2, 1, -1],
            [1, 1, -1, -2],
            [-1, -1, -2, -3],
        ]
        self.assertEqual(self.solution.countNegatives(grid), 8)

    def test_example2(self) -> None:
        grid = [
            [3, 2],
            [1, 0],
        ]
        self.assertEqual(self.solution.countNegatives(grid), 0)

    def test_all_negative(self) -> None:
        grid = [
            [-1, -1],
            [-2, -3],
        ]
        self.assertEqual(self.solution.countNegatives(grid), 4)

    def test_single_row(self) -> None:
        grid = [[5, 1, 0, -1, -2]]
        self.assertEqual(self.solution.countNegatives(grid), 2)

    def test_single_col(self) -> None:
        grid = [
            [2],
            [1],
            [0],
            [-1],
        ]
        self.assertEqual(self.solution.countNegatives(grid), 1)


if __name__ == "__main__":
    unittest.main()
