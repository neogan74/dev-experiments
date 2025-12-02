import unittest
from solution import Solution

class TestSolution(unittest.TestCase):
    def setUp(self):
        self.solution = Solution()

    def test_example_1(self):
        nums = [5, 3, 6, 1, 12]
        original = 3
        expected = 24
        result = self.solution.findFinalValue(nums, original)
        self.assertEqual(expected, result)

    def test_example_2(self):
        nums = [6, 3, 12]
        original = 2
        expected = 2
        result = self.solution.findFinalValue(nums, original)
        self.assertEqual(expected, result)

    def test_original_not_in_nums(self):
        nums = [1, 2, 3]
        original = 5
        expected = 5
        result = self.solution.findFinalValue(nums, original)
        self.assertEqual(expected, result)

    def test_multiple_doublings(self):
        nums = [4, 2, 8, 16]
        original = 2
        expected = 32
        result = self.solution.findFinalValue(nums, original)
        self.assertEqual(expected, result)

    def test_empty_nums(self):
        nums = []
        original = 10
        expected = 10
        result = self.solution.findFinalValue(nums, original)
        self.assertEqual(expected, result)

    def test_original_becomes_large(self):
        nums = [1, 2, 4, 8, 16, 32]
        original = 1
        expected = 64
        result = self.solution.findFinalValue(nums, original)
        self.assertEqual(expected, result)

if __name__ == '__main__':
    unittest.main()