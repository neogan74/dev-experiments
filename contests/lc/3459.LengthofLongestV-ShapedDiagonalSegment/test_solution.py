import unittest

from solution import Solution


class TestLenOfVDiagonal(unittest.TestCase):
    def setUp(self):
        self.solver = Solution()

    def test_examples(self):
        tests = [
            (
                "example 1",
                [
                    [2, 2, 1, 2, 2],
                    [2, 0, 2, 2, 0],
                    [2, 0, 1, 1, 0],
                    [1, 0, 2, 2, 2],
                    [2, 0, 0, 2, 2],
                ],
                5,
            ),
            (
                "example 2",
                [
                    [2, 2, 2, 2, 2],
                    [2, 0, 2, 2, 0],
                    [2, 0, 1, 1, 0],
                    [1, 0, 2, 2, 2],
                    [2, 0, 0, 2, 2],
                ],
                4,
            ),
            (
                "example 3",
                [
                    [1, 2, 2, 2, 2],
                    [2, 2, 2, 2, 0],
                    [2, 0, 0, 0, 0],
                    [0, 0, 2, 2, 2],
                    [2, 0, 0, 2, 0],
                ],
                5,
            ),
        ]

        for name, grid, expected in tests:
            with self.subTest(name=name):
                self.assertEqual(self.solver.lenOfVDiagonal(grid), expected)

    def test_single_cell_start(self):
        self.assertEqual(self.solver.lenOfVDiagonal([[1]]), 1)

    def test_no_starting_one(self):
        self.assertEqual(self.solver.lenOfVDiagonal([[2, 0], [0, 2]]), 0)

    def test_turn_required(self):
        grid = [
            [0, 0, 0, 1],
            [0, 0, 2, 0],
            [0, 0, 0, 0],
            [2, 0, 2, 2],
        ]
        # Path: (0,3) -> (1,2) -> (2,1) turn -> (3,0)
        self.assertEqual(self.solver.lenOfVDiagonal(grid), 4)


if __name__ == "__main__":
    unittest.main()
