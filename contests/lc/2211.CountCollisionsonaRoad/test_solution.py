import unittest

from solution import Solution


class TestCountCollisions(unittest.TestCase):
    def test_examples_and_edges(self):
        cases = [
            ("RLRSLL", 5),
            ("LLRR", 0),
            ("S", 0),
            ("L", 0),
            ("R", 0),
            ("SRRS", 2),
            ("RRSLL", 4),
        ]
        sol = Solution()
        for directions, expected in cases:
            with self.subTest(directions=directions):
                self.assertEqual(sol.countCollisions(directions), expected)


if __name__ == "__main__":
    unittest.main()
