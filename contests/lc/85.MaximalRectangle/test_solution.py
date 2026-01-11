import unittest

from solution import Solution


class TestMaximalRectangle(unittest.TestCase):
    def setUp(self) -> None:
        self.solver = Solution()

    def test_example1(self) -> None:
        matrix = [
            ["1", "0", "1", "0", "0"],
            ["1", "0", "1", "1", "1"],
            ["1", "1", "1", "1", "1"],
            ["1", "0", "0", "1", "0"],
        ]
        self.assertEqual(self.solver.maximalRectangle(matrix), 6)

    def test_example2(self) -> None:
        matrix = [["0"]]
        self.assertEqual(self.solver.maximalRectangle(matrix), 0)

    def test_example3(self) -> None:
        matrix = [["1"]]
        self.assertEqual(self.solver.maximalRectangle(matrix), 1)

    def test_all_zeros(self) -> None:
        matrix = [
            ["0", "0", "0"],
            ["0", "0", "0"],
        ]
        self.assertEqual(self.solver.maximalRectangle(matrix), 0)

    def test_all_ones(self) -> None:
        matrix = [
            ["1", "1"],
            ["1", "1"],
            ["1", "1"],
        ]
        self.assertEqual(self.solver.maximalRectangle(matrix), 6)

    def test_empty_matrix(self) -> None:
        self.assertEqual(self.solver.maximalRectangle([]), 0)


if __name__ == "__main__":
    unittest.main()
