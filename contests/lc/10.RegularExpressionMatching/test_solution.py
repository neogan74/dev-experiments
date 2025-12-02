import unittest
from solution import Solution

class TestRegularExpressionMatching(unittest.TestCase):
    def setUp(self):
        self.solution = Solution()

    def test_basic_matching(self):
        self.assertEqual(self.solution.isMatch("aa", "a"), False)
        self.assertEqual(self.solution.isMatch("aa", "a*"), True)
        self.assertEqual(self.solution.isMatch("ab", ".*"), True)
        self.assertEqual(self.solution.isMatch("aab", "c*a*b"), True)
        self.assertEqual(self.solution.isMatch("mississippi", "mis*is*p*."), False)

    def test_edge_cases(self):
        self.assertEqual(self.solution.isMatch("", ""), True)
        self.assertEqual(self.solution.isMatch("", "a*"), True)
        self.assertEqual(self.solution.isMatch("a", ""), False)
        self.assertEqual(self.solution.isMatch("a", ".*.."), False)

    def test_complex_patterns(self):
        self.assertEqual(self.solution.isMatch("aaa", "a.a"), True)
        self.assertEqual(self.solution.isMatch("aaa", "a*."), True)
        self.assertEqual(self.solution.isMatch("abbb", "ab*"), True)
        self.assertEqual(self.solution.isMatch("abbb", "a.c"), False)

if __name__ == '__main__':
    unittest.main()