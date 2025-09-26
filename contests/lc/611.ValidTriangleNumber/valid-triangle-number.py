from bisect import bisect_left
from typing import List

class Solution:
    def triangleNumber(self, nums: List[int]) -> int:
        nums.sort()
        ans, n = 0, len(nums)
        for i in range(n - 2):
            for j in range(i + 1, n - 1):
                k = bisect_left(nums, nums[i] + nums[j], lo=j + 1) - 1
                ans += k - j
        return ans
    
def main():
    s = Solution()
    print(s.triangleNumber([2,2,3,4]))  # Output: 3
    print(s.triangleNumber([4,2,3,4]))  # Output: 4

if __name__ == "__main__":
    main()