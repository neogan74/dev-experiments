import unittest

from solution import Solution


class TestMinNumberOperations(unittest.TestCase):
    def setUp(self) -> None:
        self.solution = Solution()

    def test_example_one(self) -> None:
        low = 3
        high = 7
        expected = 3
        self.assertEqual(self.solution.countOdds(low,high), expected)

    def test_example_one(self) -> None:
        low = 8
        high = 10
        expected = 1
        self.assertEqual(self.solution.countOdds(low,high), expected)

    def test_example_one(self) -> None:
        low = 1
        high = 101
        expected = 51
        self.assertEqual(self.solution.countOdds(low,high), expected)
        self.assertEqual(self.solution.countOdds2(low,high), expected)

    def test_example_one(self) -> None:
        low = 1
        high = 100500
        expected = 50250
        self.assertEqual(self.solution.countOdds(low,high), expected)
        self.assertEqual(self.solution.countOdds2(low,high), expected)

    def test_example_one(self) -> None:
        low = 798273637
        high = 970699661
        expected = 86213013
        self.assertEqual(self.solution.countOdds(low,high), expected)
        self.assertEqual(self.solution.countOdds2(low,high), expected)

if __name__ == "__main__":
    unittest.main()
