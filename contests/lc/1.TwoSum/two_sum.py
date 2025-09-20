from typing import List


class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        # Track indices of seen numbers to find complements in O(n)
        indices_by_value = {}
        for idx, value in enumerate(nums):
            complement = target - value
            if complement in indices_by_value:
                return [indices_by_value[complement], idx]
            indices_by_value[value] = idx
        return []

    def twoSum2(self, nums: List[int],target: int) -> List[int]:
        d = {}
        for i,x in enumerate(nums):
            if (y := target -x) in d:
                return [d[y],i]
            d[x] =i


def main() -> None:
    nums = [2, 7, 11, 15]
    target = 9
    result = Solution().twoSum(nums, target)
    result2 = Solution().twoSum2(nums,target)
    print(result,result2)


if __name__ == "__main__":
    main()
