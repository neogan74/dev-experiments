import unittest

from solution import Solution


class CountTriplesTests(unittest.TestCase):
    def test_examples(self):
        solver = Solution()
        self.assertEqual(solver.countTriples(5), 2)
        self.assertEqual(solver.countTriples(10), 4)

    def test_no_triples(self):
        solver = Solution()
        self.assertEqual(solver.countTriples(1), 0)


if __name__ == "__main__":
    unittest.main()
