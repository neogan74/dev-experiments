import random
import unittest

from successful_pairs_of_spells_and_potions import successful_pairs


def brute_successful_pairs(spells, potions, success):
    return [
        sum(1 for potion in potions if spell * potion >= success)
        for spell in spells
    ]


class SuccessfulPairsTests(unittest.TestCase):
    def test_examples(self):
        self.assertEqual(
            successful_pairs([5, 1, 3], [1, 2, 3, 4, 5], 7),
            [4, 0, 3],
        )
        self.assertEqual(
            successful_pairs([3, 1, 2], [8, 5, 8], 16),
            [2, 0, 2],
        )

    def test_edge_cases(self):
        self.assertEqual(
            successful_pairs([1, 2, 3], [1, 2, 3], 100),
            [0, 0, 0],
        )
        self.assertEqual(
            successful_pairs([10], [1], 5),
            [1],
        )

    def test_randomized(self):
        rng = random.Random(42)
        for _ in range(200):
            n = rng.randint(1, 8)
            m = rng.randint(1, 8)
            spells = [rng.randint(1, 20) for _ in range(n)]
            potions = [rng.randint(1, 20) for _ in range(m)]
            success = rng.randint(1, 400)

            self.assertEqual(
                successful_pairs(spells[:], potions[:], success),
                brute_successful_pairs(spells, potions, success),
            )


if __name__ == "__main__":
    unittest.main()
