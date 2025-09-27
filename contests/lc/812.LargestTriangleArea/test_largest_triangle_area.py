import unittest

from largest_triangle_area import largest_triangle_area


class LargestTriangleAreaTests(unittest.TestCase):
    def test_example_from_problem(self) -> None:
        points = [[0, 0], [0, 1], [1, 0], [0, 2], [2, 0]]
        self.assertAlmostEqual(largest_triangle_area(points), 2.0)

    def test_colinear_points(self) -> None:
        points = [[0, 0], [1, 1], [2, 2], [3, 3]]
        self.assertAlmostEqual(largest_triangle_area(points), 0.0)

    def test_negative_coordinates(self) -> None:
        points = [[-1, 0], [0, 0], [0, 3], [2, 0]]
        self.assertAlmostEqual(largest_triangle_area(points), 4.5)


if __name__ == "__main__":
    unittest.main()
