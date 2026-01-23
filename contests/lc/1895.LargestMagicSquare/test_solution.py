import unittest
from solution import Solution

class TestSolution(unittest.TestCase):
    def setUp(self):
        self.solution = Solution()

    def test_example_1(self):
        grid = [[7,1,4,5,6],[2,5,1,6,4],[1,5,4,3,2],[1,2,7,3,4]]
        # Explanation: 
        # The largest magic square has a size of 3.
        # One such magic square is:
        # 5 1 6 
        # 5 4 3
        # 2 7 3
        # Row sums: 5+1+6=12, 5+4+3=12, 2+7+3=12
        # Col sums: 5+5+2=12, 1+4+7=12, 6+3+3=12
        # Diagonals: 5+4+3=12, 6+4+2=12
        self.assertEqual(self.solution.largestMagicSquare(grid), 3)

    def test_example_2(self):
        grid = [[5,1,3,1],[9,3,3,1],[1,3,3,8]]
        self.assertEqual(self.solution.largestMagicSquare(grid), 2)

    def test_single_cell(self):
        grid = [[5]]
        self.assertEqual(self.solution.largestMagicSquare(grid), 1)

    def test_small_grid(self):
        grid = [[1, 2], [3, 4]]
        self.assertEqual(self.solution.largestMagicSquare(grid), 1)
        
    def test_all_same(self):
        grid = [[2, 2, 2], [2, 2, 2], [2, 2, 2]]
        self.assertEqual(self.solution.largestMagicSquare(grid), 3)

if __name__ == '__main__':
    unittest.main()
