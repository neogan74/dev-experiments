import unittest

from solution import Solution


class TestSolution(unittest.TestCase):
    def setUp(self):
        self.solver = Solution()

    def test_example_1(self):
        grid = [
            [1, 1, 1, 1],
            [2, 2, 2, 2],
            [1, 1, 1, 1],
            [2, 2, 2, 2],
        ]
        self.assertEqual(self.solver.minCost(grid), 3)

    def test_example_2(self):
        grid = [
            [1, 1, 3],
            [3, 2, 2],
            [1, 1, 4],
        ]
        self.assertEqual(self.solver.minCost(grid), 0)

    def test_example_3(self):
        grid = [
            [1, 2],
            [4, 3],
        ]
        self.assertEqual(self.solver.minCost(grid), 1)

    def test_single_cell(self):
        self.assertEqual(self.solver.minCost([[3]]), 0)

    def test_requires_changes(self):
        grid = [
            [2, 2],
            [2, 2],
        ]
        self.assertEqual(self.solver.minCost(grid), 2)


if __name__ == "__main__":
    unittest.main()
