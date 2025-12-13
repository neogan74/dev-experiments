import unittest

from solution import Solution


class TestValidCouponCodes(unittest.TestCase):
    def setUp(self) -> None:
        self.solver = Solution()

    def test_example1(self):
        code = ["SAVE20", "", "PHARMA5", "SAVE@20"]
        businessLine = ["restaurant", "grocery", "pharmacy", "restaurant"]
        isActive = [True, True, True, True]
        expected = ["PHARMA5", "SAVE20"]
        self.assertEqual(self.solver.validCouponCodes(code, businessLine, isActive), expected)

    def test_example2(self):
        code = ["GROCERY15", "ELECTRONICS_50", "DISCOUNT10"]
        businessLine = ["grocery", "electronics", "invalid"]
        isActive = [False, True, True]
        expected = ["ELECTRONICS_50"]
        self.assertEqual(self.solver.validCouponCodes(code, businessLine, isActive), expected)

    def test_sorting_and_filtering(self):
        code = ["ZPHARM", "APPLY", "B_ELEC", "AELEC", "DINNER", "BLUNCH"]
        businessLine = ["pharmacy", "restaurant", "electronics", "electronics", "restaurant", "grocery"]
        isActive = [True, True, True, True, True, True]
        expected = ["AELEC", "B_ELEC", "BLUNCH", "ZPHARM", "APPLY", "DINNER"]
        self.assertEqual(self.solver.validCouponCodes(code, businessLine, isActive), expected)

    def test_invalid_and_inactive(self):
        code = ["GOOD", "BAD-CHAR", "", "OK_UNDERSCORE"]
        businessLine = ["grocery", "grocery", "grocery", "restaurant"]
        isActive = [False, True, True, True]
        expected = ["OK_UNDERSCORE"]
        self.assertEqual(self.solver.validCouponCodes(code, businessLine, isActive), expected)


if __name__ == "__main__":
    unittest.main()
