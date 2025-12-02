import unittest

from solution import smallestRepunitDivByK


class TestSmallestRepunitDivByK(unittest.TestCase):
    def test_cases(self):
        cases = [
            (1, 1),
            (2, -1),
            (3, 3),
            (7, 6),    # 111111
            (17, 16),  # known value
            (19, 18),  # known value
            (9901, 12),
        ]
        for k, expected in cases:
            with self.subTest(k=k):
                self.assertEqual(smallestRepunitDivByK(k), expected)


if __name__ == "__main__":
    unittest.main()
