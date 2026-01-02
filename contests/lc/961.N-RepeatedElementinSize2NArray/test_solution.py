import unittest

from solution import Solution


class TestRepeatedNTimes(unittest.TestCase):
    def test_examples(self):
        solver = Solution()
        self.assertEqual(solver.repeatedNTimes([1, 2, 3, 3]), 3)
        self.assertEqual(solver.repeatedNTimes([2, 1, 2, 5, 3, 2]), 2)
        self.assertEqual(solver.repeatedNTimes([5, 1, 5, 2, 5, 3, 5, 4]), 5)

    def test_additional_cases(self):
        solver = Solution()
        self.assertEqual(solver.repeatedNTimes([9, 9, 1, 2]), 9)
        self.assertEqual(solver.repeatedNTimes([4, 1, 2, 4]), 4)


if __name__ == "__main__":
    unittest.main()
