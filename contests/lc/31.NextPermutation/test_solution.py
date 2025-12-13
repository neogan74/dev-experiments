import unittest

from solution import Solution


class NextPermutationTests(unittest.TestCase):
    def test_samples_and_edges(self):
        cases = [
            ([1, 2, 3], [1, 3, 2]),
            ([3, 2, 1], [1, 2, 3]),
            ([1, 1, 5], [1, 5, 1]),
            ([1, 3, 2], [2, 1, 3]),
            ([7], [7]),
        ]

        solver = Solution()
        for nums, expected in cases:
            with self.subTest(nums=nums):
                arr = list(nums)
                solver.nextPermutation(arr)
                self.assertEqual(expected, arr)


if __name__ == "__main__":
    unittest.main()
