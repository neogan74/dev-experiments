import unittest

from solution import prefixesDivBy5


class PrefixesDivisibleBy5Tests(unittest.TestCase):
    def test_single_zero(self):
        self.assertEqual(prefixesDivBy5([0]), [True])

    def test_example_one(self):
        self.assertEqual(prefixesDivBy5([0, 1, 1]), [True, False, False])

    def test_example_two(self):
        self.assertEqual(prefixesDivBy5([1, 1, 1]), [False, False, False])

    def test_alternating_bits(self):
        self.assertEqual(
            prefixesDivBy5([1, 0, 1, 0, 1]),
            [False, False, True, True, False],
        )

    def test_long_run_of_ones(self):
        self.assertEqual(
            prefixesDivBy5([1, 1, 1, 1, 1, 1]),
            [False, False, False, True, False, False],
        )


if __name__ == "__main__":
    unittest.main()
