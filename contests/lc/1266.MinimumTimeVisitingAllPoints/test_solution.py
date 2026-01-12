import unittest

from solution import Solution


class TestMinTimeToVisitAllPoints(unittest.TestCase):
    def setUp(self) -> None:
        self.solution = Solution()

    def test_example1(self) -> None:
        points = [[1, 1], [3, 4], [-1, 0]]
        self.assertEqual(self.solution.minTimeToVisitAllPoints(points), 7)

    def test_example2(self) -> None:
        points = [[3, 2], [-2, 2]]
        self.assertEqual(self.solution.minTimeToVisitAllPoints(points), 5)

    def test_single_point(self) -> None:
        points = [[0, 0]]
        self.assertEqual(self.solution.minTimeToVisitAllPoints(points), 0)

    def test_mixed_moves(self) -> None:
        points = [[0, 0], [1, 1], [1, 2]]
        self.assertEqual(self.solution.minTimeToVisitAllPoints(points), 2)


if __name__ == "__main__":
    unittest.main()
