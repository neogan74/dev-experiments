from typing import List

class Solution:
    def largestPerimeter(self, nums: List[int]) -> int:
        nums.sort()
        for i in range(len(nums) - 1, 1, -1):
            if (c := nums[i - 1] + nums[i - 2]) > nums[i]:
                return c + nums[i]
        return 0
    
__all__ = ["Solution"]

def main():
    import fire
    nums = [2,1,2]
    fire.Fire(Solution.largestPerimeter(nums))