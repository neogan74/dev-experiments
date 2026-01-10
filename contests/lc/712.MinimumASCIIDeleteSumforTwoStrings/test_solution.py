import unittest

from solution import Solution


def ascii_sum(value: str) -> int:
    return sum(ord(ch) for ch in value)


class TestMinimumDeleteSum(unittest.TestCase):
    def test_examples(self) -> None:
        solver = Solution()
        self.assertEqual(solver.minimumDeleteSum("sea", "eat"), 231)
        self.assertEqual(solver.minimumDeleteSum("delete", "leet"), 403)

    def test_edge_cases(self) -> None:
        solver = Solution()
        self.assertEqual(solver.minimumDeleteSum("", "abc"), ascii_sum("abc"))
        self.assertEqual(solver.minimumDeleteSum("abc", ""), ascii_sum("abc"))
        self.assertEqual(solver.minimumDeleteSum("a", "a"), 0)
        self.assertEqual(solver.minimumDeleteSum("a", "b"), ascii_sum("ab"))
        self.assertEqual(solver.minimumDeleteSum("ab", "ba"), ascii_sum("ab"))


if __name__ == "__main__":
    unittest.main()
