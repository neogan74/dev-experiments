import unittest
from solution import Solution

class TestSolution(unittest.TestCase):
    def setUp(self):
        self.solution = Solution()

    def test_example_1(self):
        bits = [1, 0, 1, 0]
        expected = True
        result = self.solution.isOneBitCharacter(bits)
        self.assertEqual(expected, result)

    def test_example_2(self):
        bits = [1, 1, 1, 0]
        expected = False
        result = self.solution.isOneBitCharacter(bits)
        self.assertEqual(expected, result)

    def test_single_bit(self):
        bits = [0]
        expected = True
        result = self.solution.isOneBitCharacter(bits)
        self.assertEqual(expected, result)

    def test_all_ones(self):
        bits = [1, 1, 1, 1, 0]
        expected = False
        result = self.solution.isOneBitCharacter(bits)
        self.assertEqual(expected, result)

    def test_no_leading_ones(self):
        bits = [0, 0, 0, 0]
        expected = True
        result = self.solution.isOneBitCharacter(bits)
        self.assertEqual(expected, result)

    def test_mixed_case(self):
        bits = [1, 0, 1, 1, 0]
        expected = True
        result = self.solution.isOneBitCharacter(bits)
        self.assertEqual(expected, result)

    def test_edge_case(self):
        bits = [1, 0, 0]
        expected = True
        result = self.solution.isOneBitCharacter(bits)
        self.assertEqual(expected, result)

if __name__ == '__main__':
    unittest.main()