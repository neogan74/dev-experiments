import unittest
from solution import Solution

class TestBestClosingTime(unittest.TestCase):

    def setUp(self):
        self.solution = Solution()
        self.bestClosingTime = self.solution.bestClosingTime

    def test_example_1(self):
        self.assertEqual(self.bestClosingTime("YYNY"), 2)

    def test_example_2(self):
        self.assertEqual(self.bestClosingTime("Y"), 1)

    def test_example_3(self):
        self.assertEqual(self.bestClosingTime("N"), 0)

    def test_all_N(self):
        self.assertEqual(self.bestClosingTime("NNNN"), 0)

    def test_all_Y(self):
        self.assertEqual(self.bestClosingTime("YYYY"), 4)

    def test_mixed(self):
        self.assertEqual(self.bestClosingTime("YNYNYN"), 1)

if __name__ == '__main__':
    unittest.main()
