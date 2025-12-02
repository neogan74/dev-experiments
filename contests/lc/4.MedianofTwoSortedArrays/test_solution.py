import unittest

from solution import Solution


class TestFindMedianSortedArrays(unittest.TestCase):
    def setUp(self):
        self.solution = Solution()

    def test_odd_total_length(self):
        result = self.solution.findMedianSortedArrays([1, 3], [2])
        self.assertEqual(result, 2.0)

    def test_even_total_length(self):
        result = self.solution.findMedianSortedArrays([1, 2], [3, 4])
        self.assertEqual(result, 2.5)

    def test_one_empty_array(self):
        result = self.solution.findMedianSortedArrays([], [1, 2, 3, 4])
        self.assertEqual(result, 2.5)

    def test_arrays_with_duplicates(self):
        result = self.solution.findMedianSortedArrays([1, 1, 1], [1, 1])
        self.assertEqual(result, 1.0)

    def test_arrays_with_negative_numbers(self):
        result = self.solution.findMedianSortedArrays([-5, -3, -1], [-2, 0, 4])
        self.assertEqual(result, -1.5)

    def test_significantly_different_lengths(self):
        result = self.solution.findMedianSortedArrays([1], [2, 3, 4, 5, 6, 7])
        self.assertEqual(result, 4.0)


if __name__ == "__main__":
    unittest.main()
