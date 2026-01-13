import unittest
from solution import Solution

class TestSeparateSquares(unittest.TestCase):
    def setUp(self):
        self.sol = Solution()

    def test_example_1(self):
        squares = [[0, 0, 1], [2, 2, 1]]
        result = self.sol.separateSquares(squares)
        self.assertAlmostEqual(result, 1.0, delta=1e-5)

    def test_example_2(self):
        squares = [[0, 0, 2], [1, 1, 1]]
        result = self.sol.separateSquares(squares)
        self.assertAlmostEqual(result, 1.16667, delta=1e-5)

    def test_single_square(self):
        squares = [[0, 0, 3]]
        result = self.sol.separateSquares(squares)
        self.assertAlmostEqual(result, 1.5, delta=1e-5)

    def test_overlapping_squares(self):
        squares = [[0, 0, 2], [0, 1, 2]]
        result = self.sol.separateSquares(squares)
        # Expected: y = 1.5
        self.assertAlmostEqual(result, 1.5, delta=1e-5)

    def test_large_input(self):
        # Many small squares
        squares = [[0, i, 1] for i in range(1000)]
        result = self.sol.separateSquares(squares)
        self.assertAlmostEqual(result, 499.5, delta=1e-5)

    def test_edge_case_all_same_y(self):
        squares = [[0, 10, 2], [5, 10, 3], [10, 10, 4]]
        result = self.sol.separateSquares(squares)
        self.assertAlmostEqual(result, 11.0, delta=1e-5)

if __name__ == '__main__':
    unittest.main()