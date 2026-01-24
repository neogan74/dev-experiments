import unittest

from solution import numOfWays


class NumOfWaysTests(unittest.TestCase):
    def test_small(self):
        self.assertEqual(numOfWays(1), 12)
        self.assertEqual(numOfWays(2), 54)
        self.assertEqual(numOfWays(3), 246)

    def test_large(self):
        self.assertEqual(numOfWays(5000), 30228214)


if __name__ == "__main__":
    unittest.main()
