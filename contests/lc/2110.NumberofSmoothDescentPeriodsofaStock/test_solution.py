import unittest

from solution import Solution
from solution import Solution2


class TestDescentPeriods(unittest.TestCase):
    def test_examples(self):
        s = Solution()
        self.assertEqual(s.getDescentPeriods([3, 2, 1, 4]), 7)
        self.assertEqual(s.getDescentPeriods([8, 6, 7, 7]), 4)
        self.assertEqual(s.getDescentPeriods([1]), 1)

    def test_additional_cases(self):
        s = Solution()
        self.assertEqual(s.getDescentPeriods([5, 4, 3, 2, 1]), 15)  # 5*6/2
        self.assertEqual(s.getDescentPeriods([5, 5, 5]), 3)
        self.assertEqual(s.getDescentPeriods([4, 3, 4, 3, 2]), 9)

    def test_examples2(self):
        s = Solution2()
        self.assertEqual(s.getDescentPeriods([3, 2, 1, 4]), 7)
        self.assertEqual(s.getDescentPeriods([8, 6, 7, 7]), 4)
        self.assertEqual(s.getDescentPeriods([1]), 1)

    def test_additional_cases(self):
        s = Solution2()
        self.assertEqual(s.getDescentPeriods([5, 4, 3, 2, 1]), 15)  # 5*6/2
        self.assertEqual(s.getDescentPeriods([5, 5, 5]), 3)
        self.assertEqual(s.getDescentPeriods([4, 3, 4, 3, 2]), 9)

if __name__ == "__main__":
    unittest.main()
