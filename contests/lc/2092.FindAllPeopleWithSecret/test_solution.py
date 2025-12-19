import unittest
from collections import defaultdict, deque
from typing import List

# Вставь здесь твой класс Solution
class Solution:
    def findAllPeople(self, n: int, meetings: List[List[int]], firstPerson: int) -> List[int]:
        vis = [False] * n
        vis[0] = vis[firstPerson] = True
        meetings.sort(key=lambda x: x[2])
        i, m = 0, len(meetings)
        while i < m:
            j = i
            while j + 1 < m and meetings[j + 1][2] == meetings[i][2]:
                j += 1
            s = set()
            g = defaultdict(list)
            for x, y, _ in meetings[i : j + 1]:
                g[x].append(y)
                g[y].append(x)
                s.update([x, y])
            q = deque([u for u in s if vis[u]])
            while q:
                u = q.popleft()
                for v in g[u]:
                    if not vis[v]:
                        vis[v] = True
                        q.append(v)
            i = j + 1
        return [i for i, v in enumerate(vis) if v]


class TestFindAllPeople(unittest.TestCase):
    def test_example_1(self):
        sol = Solution()
        n = 6
        meetings = [[1,2,5],[2,3,5],[3,4,5],[1,5,5]]
        firstPerson = 1
        expected = [0,1,2,3,4,5]
        result = sol.findAllPeople(n, meetings, firstPerson)
        self.assertEqual(result, expected)

    def test_example_2(self):
        sol = Solution()
        n = 4
        meetings = [[2,3,7],[1,3,7],[0,2,7],[0,3,7]]
        firstPerson = 3
        expected = [0,1,2,3]
        result = sol.findAllPeople(n, meetings, firstPerson)
        self.assertEqual(result, expected)

    def test_no_meetings(self):
        sol = Solution()
        n = 2
        meetings = []
        firstPerson = 1
        expected = [0,1]
        result = sol.findAllPeople(n, meetings, firstPerson)
        self.assertEqual(result, expected)

    def test_secret_not_spread(self):
        sol = Solution()
        n = 5
        meetings = [[1,2,10],[2,3,10],[1,3,11]]
        firstPerson = 4
        expected = [0,4]
        result = sol.findAllPeople(n, meetings, firstPerson)
        self.assertEqual(result, expected)

    def test_single_meeting_with_spread(self):
        sol = Solution()
        n = 3
        meetings = [[0,1,5],[1,2,5]]
        firstPerson = 0
        expected = [0,1,2]
        result = sol.findAllPeople(n, meetings, firstPerson)
        self.assertEqual(result, expected)

    def test_disconnected_components(self):
        sol = Solution()
        n = 6
        meetings = [[1,2,5],[3,4,5],[2,5,5],[0,1,5]]
        firstPerson = 1
        expected = [0,1,2,5]
        result = sol.findAllPeople(n, meetings, firstPerson)
        self.assertEqual(result, expected)


if __name__ == '__main__':
    unittest.main()