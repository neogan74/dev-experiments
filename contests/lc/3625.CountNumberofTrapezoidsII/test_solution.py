import unittest
from itertools import combinations

from solution import Solution


def _cross(a, b, c):
    return (b[0] - a[0]) * (c[1] - a[1]) - (b[1] - a[1]) * (c[0] - a[0])


def _convex_hull(points):
    pts = sorted(points)
    if len(pts) <= 1:
        return pts

    lower = []
    for p in pts:
        while len(lower) >= 2 and _cross(lower[-2], lower[-1], p) <= 0:
            lower.pop()
        lower.append(p)

    upper = []
    for p in reversed(pts):
        while len(upper) >= 2 and _cross(upper[-2], upper[-1], p) <= 0:
            upper.pop()
        upper.append(p)

    return lower[:-1] + upper[:-1]


def brute_force_trapezoids(points):
    def is_trapezoid(quad):
        hull = _convex_hull(quad)
        if len(hull) != 4:
            return False

        a, b, c, d = hull
        e0 = (b[0] - a[0], b[1] - a[1])
        e1 = (c[0] - b[0], c[1] - b[1])
        e2 = (d[0] - c[0], d[1] - c[1])
        e3 = (a[0] - d[0], a[1] - d[1])

        def parallel(u, v):
            return u[0] * v[1] - u[1] * v[0] == 0

        return parallel(e0, e2) or parallel(e1, e3)

    count = 0
    for quad in combinations(points, 4):
        if is_trapezoid(list(quad)):
            count += 1
    return count


class CountTrapezoidsTests(unittest.TestCase):
    def setUp(self):
        self.fn = Solution().countTrapezoids

    def test_examples(self):
        cases = [
            (
                [[0, 0], [2, 0], [1, 1], [3, 1]],
                1,
            ),
            (
                [[-3, 2], [3, 0], [2, 3], [3, 2], [2, -3]],
                2,
            ),
            (
                [[0, 0], [1, 0], [0, 1], [2, 1]],
                1,
            ),
        ]
        for points, expected in cases:
            with self.subTest(points=points):
                self.assertEqual(self.fn(points), expected)

    def test_no_trapezoids(self):
        cases = [
            [[0, 0], [1, 2], [2, 5], [3, 1]],
            [[0, 0], [2, 1], [3, 4], [5, 2]],
            [[0, 0], [1, 1], [2, 2], [3, 3], [0, 1]],
        ]
        for points in cases:
            with self.subTest(points=points):
                self.assertEqual(self.fn(points), 0)

    def test_matches_bruteforce_on_small_set(self):
        points = [
            [-2, 0],
            [-1, 2],
            [0, -1],
            [1, 1],
            [2, -2],
            [3, 0],
        ]
        expected = brute_force_trapezoids(points)
        self.assertEqual(self.fn(points), expected)


if __name__ == "__main__":
    unittest.main()
