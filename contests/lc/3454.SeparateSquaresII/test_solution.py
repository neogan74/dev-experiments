import unittest

from solution import Solution


class TestSeparateSquares(unittest.TestCase):
    def setUp(self):
        self.solution = Solution()

    def assertAlmostEqualAbs(self, got, want, eps=1e-6):
        self.assertLessEqual(abs(got - want), eps)

    def test_example1(self):
        squares = [[0, 0, 1], [2, 2, 1]]
        self.assertAlmostEqualAbs(self.solution.separateSquares(squares), 1.0)

    def test_example2(self):
        squares = [[0, 0, 2], [1, 1, 1]]
        self.assertAlmostEqualAbs(self.solution.separateSquares(squares), 1.0)

    def test_single_square(self):
        squares = [[5, 10, 4]]
        self.assertAlmostEqualAbs(self.solution.separateSquares(squares), 12.0)

    def test_identical_overlap(self):
        squares = [[0, 0, 2], [0, 0, 2]]
        self.assertAlmostEqualAbs(self.solution.separateSquares(squares), 1.0)

    def test_gap_between(self):
        squares = [[0, 0, 1], [0, 2, 1]]
        self.assertAlmostEqualAbs(self.solution.separateSquares(squares), 1.0)

    def test_rect_union(self):
        squares = [[0, 0, 2], [1, 0, 2]]
        self.assertAlmostEqualAbs(self.solution.separateSquares(squares), 1.0)

    def test_different_heights(self):
        squares = [[0, 0, 4], [0, 4, 2]]
        self.assertAlmostEqualAbs(self.solution.separateSquares(squares), 2.5)


if __name__ == "__main__":
    unittest.main()
