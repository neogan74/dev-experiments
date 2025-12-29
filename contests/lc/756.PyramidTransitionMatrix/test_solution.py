import unittest

from solution import Solution


class TestPyramidTransition(unittest.TestCase):
    def setUp(self) -> None:
        self.solver = Solution()

    def test_example_true(self) -> None:
        self.assertTrue(
            self.solver.pyramidTransition("BCD", ["BCC", "CDE", "CEA", "FFF"])
        )

    def test_example_false(self) -> None:
        self.assertFalse(
            self.solver.pyramidTransition("AAAA", ["AAB", "AAC", "BCD", "BBE", "DEF"])
        )

    def test_single_step(self) -> None:
        self.assertTrue(self.solver.pyramidTransition("AB", ["ABA"]))

    def test_no_rule(self) -> None:
        self.assertFalse(self.solver.pyramidTransition("AB", ["BCD"]))

    def test_branching_success(self) -> None:
        self.assertTrue(
            self.solver.pyramidTransition("ABC", ["ABD", "BCE", "BDF", "CEF"])
        )


if __name__ == "__main__":
    unittest.main()
