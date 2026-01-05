import unittest

from solution import Solution


class TestMaxMatrixSum(unittest.TestCase):
    def test_example1(self) -> None:
        matrix = [[1, -1], [-1, 1]]
        self.assertEqual(Solution().maxMatrixSum(matrix), 4)

    def test_example2(self) -> None:
        matrix = [[1, 2, 3], [-1, -2, -3], [1, 2, 3]]
        self.assertEqual(Solution().maxMatrixSum(matrix), 16)

    def test_odd_negatives_no_zero(self) -> None:
        matrix = [[-1, 2], [3, 4]]
        self.assertEqual(Solution().maxMatrixSum(matrix), 8)

    def test_odd_negatives_with_zero(self) -> None:
        matrix = [[-1, 0], [2, 3]]
        self.assertEqual(Solution().maxMatrixSum(matrix), 6)


if __name__ == "__main__":
    unittest.main()
