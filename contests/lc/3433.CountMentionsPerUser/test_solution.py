import unittest

from solution import countMentions


class CountMentionsTests(unittest.TestCase):
    def test_example1(self):
        events = [
            ["MESSAGE", "10", "id1 id0"],
            ["OFFLINE", "11", "0"],
            ["MESSAGE", "71", "HERE"],
        ]
        self.assertEqual(countMentions(2, events), [2, 2])

    def test_example2(self):
        events = [
            ["MESSAGE", "10", "id1 id0"],
            ["OFFLINE", "11", "0"],
            ["MESSAGE", "12", "ALL"],
        ]
        self.assertEqual(countMentions(2, events), [2, 2])

    def test_example3(self):
        events = [
            ["OFFLINE", "10", "0"],
            ["MESSAGE", "12", "HERE"],
        ]
        self.assertEqual(countMentions(2, events), [0, 1])

    def test_offline_same_timestamp_message(self):
        events = [
            ["OFFLINE", "5", "1"],
            ["MESSAGE", "5", "HERE"],
        ]
        self.assertEqual(countMentions(3, events), [1, 0, 1])

    def test_duplicates_and_return(self):
        events = [
            ["OFFLINE", "10", "0"],
            ["MESSAGE", "12", "id0 id0 id1"],
            ["MESSAGE", "70", "HERE"],
            ["MESSAGE", "71", "HERE"],
        ]
        self.assertEqual(countMentions(2, events), [3, 3])


if __name__ == "__main__":
    unittest.main()
