from unittest import TestCase, main
from solution import Solution

class TestNumberOfBeams(TestCase):
    def setUp(self):
        self.solution = Solution()

    def test_example_1(self):
        bank = ["011001", "000000", "010100", "001000"]
        self.assertEqual(self.solution.numberOfBeams(bank), 8)

    def test_example_2(self):
        bank = ["000", "111", "000"]
        self.assertEqual(self.solution.numberOfBeams(bank), 0)

    def test_all_devices_in_one_row(self):
        bank = ["1111", "0000", "0000"]
        self.assertEqual(self.solution.numberOfBeams(bank), 0)

    def test_two_rows_with_devices(self):
        bank = ["1100", "0011"]
        self.assertEqual(self.solution.numberOfBeams(bank), 4)

    def test_multiple_rows_no_empty(self):
        bank = ["101", "010", "101"]
        self.assertEqual(self.solution.numberOfBeams(bank), 4)

    def test_single_row(self):
        bank = ["111"]
        self.assertEqual(self.solution.numberOfBeams(bank), 0)

    def test_empty_bank(self):
        bank = []
        self.assertEqual(self.solution.numberOfBeams(bank), 0)

if __name__ == '__main__':
    main()