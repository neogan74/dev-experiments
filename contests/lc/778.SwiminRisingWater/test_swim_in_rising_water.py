import importlib.util
from pathlib import Path
import unittest


def load_solution():
    module_path = Path(__file__).with_name("swim-in-rising-water.py")
    spec = importlib.util.spec_from_file_location("swim_in_rising_water", module_path)
    module = importlib.util.module_from_spec(spec)
    spec.loader.exec_module(module)
    return module.Solution


class SwimInRisingWaterTests(unittest.TestCase):
    Solution = load_solution()

    def test_example_two_by_two(self):
        grid = [[0, 2], [1, 3]]
        self.assertEqual(self.Solution().swimInWater(grid), 3)

    def test_example_five_by_five(self):
        grid = [
            [0, 1, 2, 3, 4],
            [24, 23, 22, 21, 5],
            [12, 13, 14, 15, 16],
            [11, 17, 18, 19, 20],
            [10, 9, 8, 7, 6],
        ]
        self.assertEqual(self.Solution().swimInWater(grid), 16)

    def test_single_cell_grid(self):
        self.assertEqual(self.Solution().swimInWater([[0]]), 0)

    def test_prefers_lower_path(self):
        grid = [[0, 100], [1, 2]]
        self.assertEqual(self.Solution().swimInWater(grid), 2)

    def test_sparse_high_values_off_path(self):
        grid = [
            [0, 1, 2],
            [12, 13, 3],
            [11, 14, 4],
        ]
        self.assertEqual(self.Solution().swimInWater(grid), 4)


if __name__ == "__main__":
    unittest.main()
