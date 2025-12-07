class Solution:
    def countOdds(self, low: int, high: int) -> int:
        count = 0
        for i in range(low, high + 1):
            if i % 2 != 0:
                count += 1
        return count
    

if __name__ == "__main__":
    sol = Solution()
    print(sol.countOdds(3, 7))  # Output: 3
    print(sol.countOdds(8, 10)) # Output: 1