import unittest
from typing import List

class TestMagicSquares(unittest.TestCase):
    def setUp(self):
        self.solution = Solution()

    def test_example_1(self):
        grid = [[4,3,8,4],[9,5,1,9],[2,7,6,2]]
        self.assertEqual(self.solution.numMagicSquaresInside(grid), 1)

    def test_example_2(self):
        grid = [[8]]
        self.assertEqual(self.solution.numMagicSquaresInside(grid), 0)

    def test_small_grid(self):
        grid = [[4,3],[9,5]]
        self.assertEqual(self.solution.numMagicSquaresInside(grid), 0)

    def test_valid_magic_square(self):
        grid = [[2,7,6],[9,5,1],[4,3,8]]
        self.assertEqual(self.solution.numMagicSquaresInside(grid), 1)

    def test_invalid_magic_square(self):
        grid = [[2,7,6],[9,5,1],[4,3,9]]  # Contains duplicate 9
        self.assertEqual(self.solution.numMagicSquaresInside(grid), 0)

    def test_grid_with_multiple_candidates(self):
        grid = [
            [4,3,8,4,2],
            [9,5,1,9,7],
            [2,7,6,2,8],
            [4,3,8,4,2],
            [9,5,1,9,7]
        ]
        self.assertEqual(self.solution.numMagicSquaresInside(grid), 2)

if __name__ == '__main__':
    unittest.main()