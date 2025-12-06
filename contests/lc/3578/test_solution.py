import unittest

from solution import Solution


class TestCountPartitions(unittest.TestCase):
    def setUp(self):
        self.sol = Solution()

    def test_example_one(self):
        nums = [9, 4, 1, 3, 7]
        self.assertEqual(self.sol.countPartitions(nums, 4), 6)

    def test_example_two(self):
        nums = [3, 3, 4]
        self.assertEqual(self.sol.countPartitions(nums, 0), 2)

    def test_all_partitions_valid(self):
        nums = [1, 5, 2]
        # k is large, so every partition is allowed: 2^(n-1) = 4
        self.assertEqual(self.sol.countPartitions(nums, 10), 4)

    def test_strict_window(self):
        nums = [1, 3, 2]
        # Valid partitions: [1][3][2], [1][3,2]
        self.assertEqual(self.sol.countPartitions(nums, 1), 2)

    def test_all_equal_k_zero(self):
        nums = [5, 5, 5]
        # All partitions valid: 2^(n-1) = 4
        self.assertEqual(self.sol.countPartitions(nums, 0), 4)


if __name__ == "__main__":
    unittest.main()
