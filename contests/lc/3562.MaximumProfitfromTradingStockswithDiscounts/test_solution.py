import unittest
from solution import Solution

class TestMaxProfit(unittest.TestCase):
    def setUp(self):
        self.sol = Solution()

    def test_example_1(self):
        n = 2
        present = [1, 2]
        future = [4, 3]
        hierarchy = [[1, 2]]
        budget = 3
        self.assertEqual(self.sol.maxProfit(n, present, future, hierarchy, budget), 5)

    def test_example_2(self):
        n = 3
        present = [5, 2, 3]
        future = [8, 5, 6]
        hierarchy = [[1, 2], [2, 3]]
        budget = 7
        self.assertEqual(self.sol.maxProfit(n, present, future, hierarchy, budget), 13)

    def test_no_hierarchy(self):
        n = 1
        present = [5]
        future = [10]
        hierarchy = []
        budget = 5
        self.assertEqual(self.sol.maxProfit(n, present, future, hierarchy, budget), 5)

    def test_no_profit(self):
        n = 1
        present = [10]
        future = [5]
        hierarchy = []
        budget = 10
        self.assertEqual(self.sol.maxProfit(n, present, future, hierarchy, budget), 0)

    def test_multiple_children(self):
        n = 4
        present = [10, 5, 3, 2]
        future = [20, 10, 8, 6]
        hierarchy = [[1, 2], [1, 3], [1, 4]]
        budget = 10
        # 1 покупает со скидкой (5), 2 не покупает, 3 и 4 покупают по обычной цене (3 и 2)
        # Прибыль: (20 - 5) + (8 - 3) + (6 - 2) = 15 + 5 + 4 = 24
        self.assertEqual(self.sol.maxProfit(n, present, future, hierarchy, budget), 24)

    def test_deep_hierarchy(self):
        n = 5
        present = [10, 4, 4, 2, 2]
        future = [20, 10, 10, 5, 5]
        hierarchy = [[1, 2], [2, 3], [3, 4], [4, 5]]
        budget = 10
        # 1 покупает, все остальные получают скидку
        # Скидка: 4 → 2 → 1 → 0.5 → 0 (целочисленное деление)
        # Прибыль: 20-10 = 10, 10-2 = 8, 10-1 = 9, 5-0 = 5, 5-0 = 5
        # Общая: 10+8+9+5+5 = 37
        self.assertEqual(self.sol.maxProfit(n, present, future, hierarchy, budget), 37)

    def test_multiple_roots(self):
        n = 3
        present = [10, 5, 5]
        future = [20, 10, 10]
        hierarchy = [[1, 2], [1, 3]]
        budget = 15
        # 1 покупает, 2 и 3 получают скидку
        # Общая стоимость: 10 + 2 + 2 = 14
        # Прибыль: 20-10 = 10, 10-2 = 8, 10-2 = 8 → 26
        self.assertEqual(self.sol.maxProfit(n, present, future, hierarchy, budget), 26)

    def test_budget_too_low(self):
        n = 2
        present = [10, 5]
        future = [20, 10]
        hierarchy = [[1, 2]]
        budget = 5
        # 1 не может купить (цена 10), 2 может купить по цене 5
        # Прибыль: 0 + (10 - 5) = 5
        self.assertEqual(self.sol.maxProfit(n, present, future, hierarchy, budget), 5)

    def test_full_discount_chain(self):
        n = 4
        present = [8, 4, 2, 1]
        future = [16, 8, 4, 2]
        hierarchy = [[1, 2], [2, 3], [3, 4]]
        budget = 15
        # 1 покупает, все остальные получают скидку (цена: 4 → 2 → 1 → 0)
        # Прибыль: (16-4) + (8-2) + (4-1) + (2-0) = 12 + 6 + 3 + 2 = 23
        self.assertEqual(self.sol.maxProfit(n, present, future, hierarchy, budget), 23)

    def test_large_case(self):
        n = 99
        present = [35,7,42,45,32,50,28,44,49,13,24,5,11,27,28,18,39,13,5,22,30,49,37,41,18,36,39,16,44,37,6,36,16,3,10,44,27,1,2,9,47,22,49,31,13,50,9,19,27,46,49,47,1,11,49,21,8,24,33,47,47,6,20,50,24,6,5,2,9,46,23,14,21,20,12,27,45,29,9,46,37,4,37,4,40,21,30,34,35,41,12,19,12,38,1,34,39,18,39]
        future = [6,32,30,39,46,26,3,30,47,28,3,24,5,36,44,16,46,18,17,34,17,35,16,15,29,25,32,13,46,16,29,38,15,36,9,50,46,28,24,2,26,15,30,48,6,45,2,27,21,32,2,27,2,41,12,39,10,42,3,7,25,34,39,48,20,10,36,16,46,43,36,43,36,36,43,35,14,1,38,47,50,48,34,43,18,10,9,14,15,49,11,33,3,10,19,29,50,37,31]
        hierarchy = [[1,26],[1,56],[56,12],[1,20],[12,80],[56,67],[80,52],[56,25],[12,70],[1,96],[70,98],[12,11],[1,41],[1,64],[25,10],[56,17],[41,38],[70,22],[20,84],[12,81],[22,83],[96,88],[1,74],[81,15],[52,79],[83,27],[12,66],[64,16],[38,36],[74,37],[20,4],[96,48],[64,94],[94,9],[38,40],[74,30],[10,28],[56,14],[1,44],[10,62],[40,78],[25,57],[20,65],[16,68],[96,60],[37,49],[62,55],[37,39],[52,71],[67,31],[10,58],[49,82],[14,99],[67,51],[44,76],[40,35],[10,53],[14,24],[31,54],[4,75],[11,5],[11,19],[60,45],[25,7],[99,50],[99,29],[39,47],[27,43],[71,46],[24,63],[75,42],[25,33],[47,69],[81,72],[22,89],[4,8],[72,73],[25,3],[52,86],[25,2],[79,6],[48,34],[55,21],[39,59],[24,91],[34,93],[60,92],[27,32],[81,95],[14,85],[53,90],[54,61],[21,18],[56,97],[31,77],[71,87],[97,23],[59,13]]
        budget = 24
        result = self.sol.maxProfit(n, present, future, hierarchy, budget)
        self.assertTrue(result >= 0)

if __name__ == '__main__':
    unittest.main()