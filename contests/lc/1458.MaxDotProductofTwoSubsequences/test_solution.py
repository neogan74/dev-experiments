import unittest

from solution import maxDotProduct


class TestMaxDotProduct(unittest.TestCase):
    def test_examples(self) -> None:
        self.assertEqual(maxDotProduct([2, 1, -2, 5], [3, 0, -6]), 18)
        self.assertEqual(maxDotProduct([3, -2], [2, -6, 7]), 21)
        self.assertEqual(maxDotProduct([-1, -1], [1, 1]), -1)

    def test_single_negative_best(self) -> None:
        self.assertEqual(maxDotProduct([-5, -1], [2, 3]), -2)

    def test_single_pair(self) -> None:
        self.assertEqual(maxDotProduct([7], [-3]), -21)


if __name__ == "__main__":
    unittest.main()
