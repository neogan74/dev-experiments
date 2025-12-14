import unittest

from solution import numberOfWays


class NumberOfWaysTests(unittest.TestCase):
    def test_samples(self):
        self.assertEqual(numberOfWays("SSPPSPS"), 3)
        self.assertEqual(numberOfWays("PPSPSP"), 1)
        self.assertEqual(numberOfWays("S"), 0)

    def test_additional(self):
        cases = [
            ("SSPPSS", 3),
            ("SPPSPSSPPS", 0),
            ("PPP", 0),
            ("SS", 1),
            ("PSSPPSPSPPSSPSP", 0),
        ]
        for corridor, expected in cases:
            with self.subTest(corridor=corridor):
                self.assertEqual(numberOfWays(corridor), expected)


if __name__ == "__main__":
    unittest.main()
