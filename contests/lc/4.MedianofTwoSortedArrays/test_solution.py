import unittest

from solution import Solution


class TestMedianOfTwoSortedArrays(unittest.TestCase):
    def setUp(self):
        self.solution = Solution()

    def test_single_non_empty_array(self):
        self.assertEqual(self.solution.findMedianSortedArrays([42], []), 42.0)
        self.assertEqual(self.solution.findMedianSortedArrays([], [3]), 3.0)

    def test_leetcode_examples(self):
        self.assertEqual(self.solution.findMedianSortedArrays([1, 3], [2]), 2.0)
        self.assertEqual(
            self.solution.findMedianSortedArrays([1, 2], [3, 4]),
            2.5,
        )

    def test_handles_negative_numbers(self):
        self.assertEqual(
            self.solution.findMedianSortedArrays([-5, -3], [-2, -1]),
            -2.5,
        )

    def test_handles_duplicate_values(self):
        self.assertEqual(
            self.solution.findMedianSortedArrays([1, 1, 1], [1, 1]),
            1.0,
        )

    def test_mixed_sizes(self):
        self.assertEqual(
            self.solution.findMedianSortedArrays([1, 2, 3, 4, 5], [6, 7, 8]),
            4.5,
        )


if __name__ == "__main__":
    unittest.main()
