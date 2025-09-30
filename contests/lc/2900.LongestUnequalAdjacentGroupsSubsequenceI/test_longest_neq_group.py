import unittest

from longest_neq_group import Solution


class TestGetLongestSubsequence(unittest.TestCase):
    def setUp(self) -> None:
        self.solution = Solution()

    def test_single_element_zero_group(self) -> None:
        words = ["c"]
        groups = [0]
        self.assertEqual(self.solution.getLongestSubsequence(words, groups), ["c"])

    def test_single_element_one_group(self) -> None:
        words = ["d"]
        groups = [1]
        self.assertEqual(self.solution.getLongestSubsequence(words, groups), ["d"])

    def test_example_one(self) -> None:
        words = ["e", "a", "b"]
        groups = [0, 0, 1]
        self.assertEqual(self.solution.getLongestSubsequence(words, groups), ["e", "b"])

    def test_example_two(self) -> None:
        words = ["a", "b", "c", "d"]
        groups = [1, 0, 1, 1]
        self.assertEqual(self.solution.getLongestSubsequence(words, groups), ["a", "b", "c"])

    def test_longer_alternating(self) -> None:
        words = ["w", "x", "y", "z", "u"]
        groups = [0, 1, 0, 1, 0]
        self.assertEqual(self.solution.getLongestSubsequence(words, groups), words)

    def test_repeated_groups(self) -> None:
        words = ["aa", "bb", "cc", "dd"]
        groups = [1, 1, 0, 0]
        self.assertEqual(self.solution.getLongestSubsequence(words, groups), ["aa", "cc"])


if __name__ == "__main__":
    unittest.main()
