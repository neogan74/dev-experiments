import unittest
from solution import Solution

class TestMinDeletionSize(unittest.TestCase):
    """Unit tests for the min_deletion_size function."""
    def setUp(self):
        self.sol = Solution()

    def test_example1(self):
        self.assertEqual(self.sol.min_deletion_size(["cba", "daf", "ghi"]), 1)

    def test_example2(self):
        self.assertEqual(self.sol.min_deletion_size(["abc", "bce", "cae"]), 1)

    def test_example3(self):
        self.assertEqual(self.sol.min_deletion_size(["abc", "acb", "bac"]), 2)

    def test_example4(self):
        self.assertEqual(self.sol.min_deletion_size(["a", "b", "c"]), 0)

    def test_example5(self):
        self.assertEqual(self.sol.min_deletion_size(["abcde"]), 0)

    def test_empty_input(self):
        self.assertEqual(self.sol.min_deletion_size([]), 0)

    def test_single_column(self):
        self.assertEqual(self.sol.min_deletion_size(["a", "b", "c", "d"]), 0)

    def test_all_columns_unsorted(self):
        self.assertEqual(self.sol.min_deletion_size(["bac", "acb", "cba"]), 3)

if __name__ == '__main__':
    unittest.main()