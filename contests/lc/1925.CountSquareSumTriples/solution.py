class Solution:
    def countTriples(self, n: int) -> int:
        squares = {i * i for i in range(1, n + 1)}
        limit = n * n
        count = 0

        for a in range(1, n + 1):
            aa = a * a
            for b in range(1, n + 1):
                s = aa + b * b
                if s > limit:
                    break  # c would be larger than n as b grows
                if s in squares:
                    count += 1

        return count
