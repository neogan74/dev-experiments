import importlib.util
from pathlib import Path
from typing import List
import unittest


def load_solution():
    module_path = Path(__file__).resolve().parent / "largest-perimeter-triange.py"
    module_name = "largest_perimeter_triangle_solution"
    spec = importlib.util.spec_from_file_location(module_name, module_path)
    module = importlib.util.module_from_spec(spec)
    module.__dict__["List"] = List
    spec.loader.exec_module(module)
    return module.Solution()


class LargestPerimeterTriangleTests(unittest.TestCase):
    @classmethod
    def setUpClass(cls):
        cls.solution = load_solution()

    def test_returns_perimeter_for_valid_triangle(self):
        self.assertEqual(self.solution.largestPerimeter([2, 1, 2]), 5)

    def test_returns_zero_when_no_triangle_possible(self):
        self.assertEqual(self.solution.largestPerimeter([1, 2, 1, 10]), 0)

    def test_ignores_non_qualifying_largest_elements(self):
        # The two largest values cannot form a triangle with the third largest, so the next set should be used.
        self.assertEqual(self.solution.largestPerimeter([3, 6, 2, 3]), 8)

    def test_handles_multiple_identical_sides(self):
        self.assertEqual(self.solution.largestPerimeter([6, 6, 6]), 18)


if __name__ == "__main__":
    unittest.main()
