import unittest

from solution import Solution


class TestSolution(unittest.TestCase):
    def test_examples(self) -> None:
        solver = Solution()
        self.assertEqual(
            solver.maximizeSquareHoleArea(2, 1, [2, 3], [2]),
            4,
        )
        self.assertEqual(
            solver.maximizeSquareHoleArea(1, 1, [2], [2]),
            4,
        )
        self.assertEqual(
            solver.maximizeSquareHoleArea(2, 3, [2, 3], [2, 4]),
            4,
        )

    def test_non_consecutive(self) -> None:
        solver = Solution()
        self.assertEqual(
            solver.maximizeSquareHoleArea(5, 5, [2, 4, 6], [3, 5]),
            4,
        )

    def test_longer_run(self) -> None:
        solver = Solution()
        self.assertEqual(
            solver.maximizeSquareHoleArea(5, 5, [2, 3, 4], [2, 4, 5]),
            9,
        )


if __name__ == "__main__":
    unittest.main()
