import unittest

from solution import Solution


class TestMaxKDivisibleComponents(unittest.TestCase):
    def test_cases(self):
        cases = [
            dict(
                name="example 1",
                n=5,
                edges=[[0, 2], [1, 2], [1, 3], [2, 4]],
                values=[1, 8, 1, 4, 4],
                k=6,
                want=2,
            ),
            dict(
                name="example 2",
                n=7,
                edges=[[0, 1], [0, 2], [1, 3], [1, 4], [2, 5], [2, 6]],
                values=[3, 0, 6, 1, 5, 2, 1],
                k=3,
                want=3,
            ),
            dict(
                name="all nodes divisible when k is one",
                n=4,
                edges=[[0, 1], [1, 2], [2, 3]],
                values=[5, 1, 2, 3],
                k=1,
                want=4,
            ),
            dict(
                name="single node",
                n=1,
                edges=[],
                values=[6],
                k=3,
                want=1,
            ),
            dict(
                name="no removable edge still valid",
                n=2,
                edges=[[0, 1]],
                values=[2, 3],
                k=5,
                want=1,
            ),
        ]

        for case in cases:
            with self.subTest(case["name"]):
                got = Solution().maxKDivisibleComponents(
                    case["n"], case["edges"], case["values"], case["k"]
                )
                self.assertEqual(got, case["want"])


if __name__ == "__main__":
    unittest.main()
