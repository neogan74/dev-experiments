from typing import List


class Solution:
    """Provides hash-map based solutions for the Two Sum problem."""

    def twoSum(self, nums: List[int], target: int) -> List[int]:
        """Return indices of the two numbers that add up to ``target``.

        Uses a single-pass dictionary to track complements for O(n) time.
        """

        # Track indices of seen numbers to find complements in O(n)
        indices_by_value = {}
        for idx, value in enumerate(nums):
            complement = target - value
            if complement in indices_by_value:
                return [indices_by_value[complement], idx]
            indices_by_value[value] = idx
        return []

    def twoSum2(self, nums: List[int], target: int) -> List[int]:
        """Variant implementation mirroring ``twoSum`` using assignment expressions."""

        seen_indices = {}
        for idx, value in enumerate(nums):
            if (complement := target - value) in seen_indices:
                return [seen_indices[complement], idx]
            seen_indices[value] = idx


def main() -> None:
    nums = [2, 7, 11, 15]
    target = 9
    result = Solution().twoSum(nums, target)
    result2 = Solution().twoSum2(nums, target)
    print(result, result2)


if __name__ == "__main__":
    main()
