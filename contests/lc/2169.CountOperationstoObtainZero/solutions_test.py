import unittest
from solutions import Solution


class TestSolution(unittest.TestCase):
    def setUp(self):
        self.solution = Solution()

    def test_example1(self):
        """Example 1: num1=2, num2=3"""
        self.assertEqual(self.solution.countOperations(2, 3), 3)

    def test_example2(self):
        """Example 2: num1=10, num2=10"""
        self.assertEqual(self.solution.countOperations(10, 10), 1)

    def test_both_zero(self):
        """Both zero"""
        self.assertEqual(self.solution.countOperations(0, 0), 0)

    def test_first_zero(self):
        """First zero"""
        self.assertEqual(self.solution.countOperations(0, 5), 0)

    def test_second_zero(self):
        """Second zero"""
        self.assertEqual(self.solution.countOperations(5, 0), 0)

    def test_num1_greater(self):
        """num1 > num2"""
        self.assertEqual(self.solution.countOperations(7, 3), 5)

    def test_num2_greater(self):
        """num2 > num1"""
        self.assertEqual(self.solution.countOperations(3, 7), 5)

    def test_both_one(self):
        """Both are 1"""
        self.assertEqual(self.solution.countOperations(1, 1), 1)

    def test_large_numbers(self):
        """Large numbers"""
        self.assertEqual(self.solution.countOperations(100, 10), 10)


if __name__ == '__main__':
    unittest.main()
