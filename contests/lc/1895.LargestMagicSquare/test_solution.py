import unittest

from solution import Solution


class TestLargestMagicSquare(unittest.TestCase):
    def setUp(self) -> None:
        self.solver = Solution()

    def test_example_1(self) -> None:
        grid = [
            [7, 1, 4, 5, 6],
            [2, 5, 1, 6, 4],
            [1, 5, 4, 3, 2],
            [1, 2, 7, 3, 4],
        ]
        self.assertEqual(self.solver.largestMagicSquare(grid), 3)

    def test_example_2(self) -> None:
        grid = [
            [5, 1, 3, 1],
            [9, 3, 3, 1],
            [1, 3, 3, 8],
        ]
        self.assertEqual(self.solver.largestMagicSquare(grid), 2)

    def test_single_cell(self) -> None:
        grid = [[42]]
        self.assertEqual(self.solver.largestMagicSquare(grid), 1)

    def test_no_large_magic_square(self) -> None:
        grid = [
            [1, 2, 3],
            [4, 5, 6],
            [7, 8, 9],
        ]
        self.assertEqual(self.solver.largestMagicSquare(grid), 1)

    def test_rectangular(self) -> None:
        grid = [
            [2, 7, 6, 2],
            [9, 5, 1, 9],
            [4, 3, 8, 4],
        ]
        self.assertEqual(self.solver.largestMagicSquare(grid), 3)


if __name__ == "__main__":
    unittest.main()
