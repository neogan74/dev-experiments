import unittest

import solution


class TestCountCoveredBuildings(unittest.TestCase):
    def setUp(self) -> None:
        self.sol = solution.Solution()

    def test_examples(self):
        cases = [
            (3, [[1, 2], [2, 2], [3, 2], [2, 1], [2, 3]], 1),
            (3, [[1, 1], [1, 2], [2, 1], [2, 2]], 0),
            (5, [[1, 3], [3, 2], [3, 3], [3, 5], [5, 3]], 1),
        ]
        for n, buildings, expected in cases:
            with self.subTest(n=n, buildings=buildings):
                self.assertEqual(self.sol.countCoveredBuildings(n, buildings), expected)

    def test_full_grid_multiple_covered(self):
        n = 4
        buildings = [[r, c] for r in range(1, n + 1) for c in range(1, n + 1)]
        # Only interior cells (2,2), (2,3), (3,2), (3,3) are covered.
        self.assertEqual(self.sol.countCoveredBuildings(n, buildings), 4)

    def test_sparse_no_coverage(self):
        n = 4
        buildings = [[2, 2], [2, 4], [4, 2], [4, 4]]
        self.assertEqual(self.sol.countCoveredBuildings(n, buildings), 0)


if __name__ == "__main__":  # pragma: no cover
    unittest.main()
