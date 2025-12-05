import unittest

from solution import countPartitions


class CountPartitionsTests(unittest.TestCase):
    def test_examples(self):
        cases = [
            ([10, 10, 3, 7, 6], 4),
            ([1, 2, 2], 0),
            ([2, 4, 6, 8], 3),
        ]
        for nums, expected in cases:
            with self.subTest(nums=nums):
                self.assertEqual(countPartitions(nums), expected)


if __name__ == "__main__":
    unittest.main()
