import unittest

from water_bottle_2 import Solution, maxBottlesDrunk


class MaxBottlesDrunkTests(unittest.TestCase):
    def test_examples_match_reference_answers(self) -> None:
        cases = (
            (13, 6, 15),
            (10, 3, 13),
            (2, 3, 2),
        )
        for num_bottles, num_exchange, expected in cases:
            with self.subTest(num_bottles=num_bottles, num_exchange=num_exchange):
                self.assertEqual(maxBottlesDrunk(num_bottles, num_exchange), expected)

    def test_solution_wrapper_delegates_to_helper(self) -> None:
        self.assertEqual(Solution().maxBottlesDrunk(9, 3), maxBottlesDrunk(9, 3))


if __name__ == "__main__":
    unittest.main()
