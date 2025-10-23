import unittest

from solution import max_distinct_elements


class MaxDistinctElementsTests(unittest.TestCase):
    def test_examples(self) -> None:
        self.assertEqual(max_distinct_elements([1, 2, 2, 3, 3, 4], 2), 6)
        self.assertEqual(max_distinct_elements([4, 4, 4, 4], 1), 3)

    def test_single_element(self) -> None:
        self.assertEqual(max_distinct_elements([7], 0), 1)

    def test_no_extra_moves(self) -> None:
        self.assertEqual(max_distinct_elements([1, 1, 2, 2], 0), 2)

    def test_large_k_spread(self) -> None:
        self.assertEqual(max_distinct_elements([5, 5, 5], 5), 3)


if __name__ == "__main__":
    unittest.main()
