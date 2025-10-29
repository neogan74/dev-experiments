import unittest

from solution import Bank


class BankTests(unittest.TestCase):
    def test_operations(self) -> None:
        bank = Bank([10, 100, 20, 50, 30])

        self.assertTrue(bank.withdraw(3, 10))
        self.assertTrue(bank.transfer(5, 1, 20))
        self.assertTrue(bank.deposit(5, 20))
        self.assertFalse(bank.transfer(3, 4, 15))
        self.assertFalse(bank.withdraw(10, 50))

        self.assertEqual(bank.balance, [30, 100, 10, 50, 30])

    def test_constructor_copies_input(self) -> None:
        data = [1, 2, 3]
        bank = Bank(data)
        data[0] = 999
        self.assertEqual(bank.balance[0], 1)

    def test_edge_cases(self) -> None:
        bank = Bank([100])
        self.assertTrue(bank.deposit(1, 0))
        self.assertTrue(bank.withdraw(1, 0))
        self.assertTrue(bank.transfer(1, 1, 50))
        self.assertEqual(bank.balance, [100])
        self.assertFalse(bank.deposit(2, 10))


if __name__ == "__main__":
    unittest.main()
