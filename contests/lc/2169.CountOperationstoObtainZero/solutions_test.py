import pytest
from solutions import Solution


class TestSolution:
    def setup_method(self):
        self.solution = Solution()

    def test_example1(self):
        """Example 1: num1=2, num2=3"""
        assert self.solution.countOperations(2, 3) == 3

    def test_example2(self):
        """Example 2: num1=10, num2=10"""
        assert self.solution.countOperations(10, 10) == 1

    def test_both_zero(self):
        """Both zero"""
        assert self.solution.countOperations(0, 0) == 0

    def test_first_zero(self):
        """First zero"""
        assert self.solution.countOperations(0, 5) == 0

    def test_second_zero(self):
        """Second zero"""
        assert self.solution.countOperations(5, 0) == 0

    def test_num1_greater(self):
        """num1 > num2"""
        assert self.solution.countOperations(7, 3) == 5

    def test_num2_greater(self):
        """num2 > num1"""
        assert self.solution.countOperations(3, 7) == 5

    def test_both_one(self):
        """Both are 1"""
        assert self.solution.countOperations(1, 1) == 1

    def test_large_numbers(self):
        """Large numbers"""
        assert self.solution.countOperations(100, 10) == 10
