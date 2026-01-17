import unittest

from solution import Solution


class TestLargestSquareArea(unittest.TestCase):
    def setUp(self) -> None:
        self.solution = Solution()

    def test_example_1(self) -> None:
        bottom_left = [[1, 1], [2, 2], [3, 1]]
        top_right = [[3, 3], [4, 4], [6, 6]]
        self.assertEqual(self.solution.largestSquareArea(bottom_left, top_right), 1)

    def test_example_2(self) -> None:
        bottom_left = [[1, 1], [1, 3], [1, 5]]
        top_right = [[5, 5], [5, 7], [5, 9]]
        self.assertEqual(self.solution.largestSquareArea(bottom_left, top_right), 4)

    def test_example_3(self) -> None:
        bottom_left = [[1, 1], [2, 2], [1, 2]]
        top_right = [[3, 3], [4, 4], [3, 4]]
        self.assertEqual(self.solution.largestSquareArea(bottom_left, top_right), 1)

    def test_example_4(self) -> None:
        bottom_left = [[1, 1], [3, 3], [3, 1]]
        top_right = [[2, 2], [4, 4], [4, 2]]
        self.assertEqual(self.solution.largestSquareArea(bottom_left, top_right), 0)

    def test_overlap(self) -> None:
        bottom_left = [[1, 1], [1, 1]]
        top_right = [[6, 4], [4, 8]]
        self.assertEqual(self.solution.largestSquareArea(bottom_left, top_right), 9)


if __name__ == "__main__":
    unittest.main()
