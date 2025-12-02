import unittest

from solution import Solution


class TestMinNumberOperations(unittest.TestCase):
    def setUp(self) -> None:
        self.solution = Solution()

    def test_example_one(self) -> None:
        target = [1, 2, 3, 2, 1]
        expected = 3
        self.assertEqual(self.solution.minNumberOperations(target), expected)

    def test_example_two(self) -> None:
        target = [3, 1, 1, 2]
        expected = 4
        self.assertEqual(self.solution.minNumberOperations(target), expected)

    def test_example_three(self) -> None:
        target = [3, 1, 5, 4, 2]
        expected = 7
        self.assertEqual(self.solution.minNumberOperations(target), expected)

    def test_single_element(self) -> None:
        target = [5]
        expected = 5
        self.assertEqual(self.solution.minNumberOperations(target), expected)

    def test_monotonic_increase(self) -> None:
        target = [1, 3, 6, 10]
        expected = 10
        self.assertEqual(self.solution.minNumberOperations(target), expected)

    def test_monotonic_decrease(self) -> None:
        target = [7, 5, 3, 1]
        expected = 7
        self.assertEqual(self.solution.minNumberOperations(target), expected)

    def test_mixed_fluctuations(self) -> None:
        target = [0, 2, 2, 5, 1, 1, 4]
        expected = 8
        self.assertEqual(self.solution.minNumberOperations(target), expected)


if __name__ == "__main__":
    unittest.main()
