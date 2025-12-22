import unittest
from solution import Solution

class TestCombinationSum2(unittest.TestCase):
    def setUp(self):
        self.sol = Solution()

    def test_case_1(self):
        candidates = [10, 1, 2, 7, 6, 1, 5]
        target = 8
        expected = [
            [1, 1, 6],
            [1, 2, 5],
            [1, 7],
            [2, 6]
        ]
        result = self.sol.combinationSum2(candidates, target)
        self.assertEqual(sorted(result), sorted(expected))

    def test_case_2(self):
        candidates = [2, 5, 2, 1, 2]
        target = 5
        expected = [
            [1, 2, 2],
            [5]
        ]
        result = self.sol.combinationSum2(candidates, target)
        self.assertEqual(sorted(result), sorted(expected))

    def test_no_combinations(self):
        candidates = [2, 3, 6, 7]
        target = 1
        expected = []
        result = self.sol.combinationSum2(candidates, target)
        self.assertEqual(result, expected)

    def test_empty_candidates(self):
        candidates = []
        target = 5
        expected = []
        result = self.sol.combinationSum2(candidates, target)
        self.assertEqual(result, expected)

    def test_duplicate_handling(self):
        candidates = [1, 1, 1, 1]
        target = 2
        expected = [[1, 1]]
        result = self.sol.combinationSum2(candidates, target)
        self.assertEqual(result, expected)

if __name__ == '__main__':
    unittest.main()