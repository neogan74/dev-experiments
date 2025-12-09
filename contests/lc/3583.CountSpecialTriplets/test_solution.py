import unittest

from solution import count_special_triplets


class TestCountSpecialTriplets(unittest.TestCase):
    def test_examples_function(self):
        self.assertEqual(count_special_triplets([6, 3, 6]), 1)
        self.assertEqual(count_special_triplets([0, 1, 0, 0]), 1)
        self.assertEqual(count_special_triplets([8, 4, 2, 8, 4]), 2)

    def test_examples_method(self):
        from solution import Solution

        solver = Solution()
        self.assertEqual(solver.count_special_triplets([6, 3, 6]), 1)
        self.assertEqual(solver.specialTriplets([0, 1, 0, 0]), 1)
        self.assertEqual(solver.specialTriplets([8, 4, 2, 8, 4]), 2)

    def test_no_triplets_function(self):
        self.assertEqual(count_special_triplets([1, 2, 3, 4]), 0)
        self.assertEqual(count_special_triplets([2, 2, 2]), 0)

    def test_multiple_same_targets_function(self):
        # Middle value 3 has two doubles on the left and two on the right.
        self.assertEqual(count_special_triplets([6, 6, 3, 6, 6]), 4)

    def test_benchmark_placeholders(self):
        """
        Lightweight loop to mimic a micro-benchmark. Python's unittest
        doesn't have built-in benchmarking, but this keeps parity with Go benchmarks.
        """
        from solution import Solution

        nums = [8, 4, 2, 8, 4, 16, 8, 4, 2, 1, 2, 4, 8, 16, 32]
        solver = Solution()
        for _ in range(1000):
            count_special_triplets(nums)
            solver.specialTriplets(nums)


if __name__ == "__main__":
    unittest.main()
