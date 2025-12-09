from collections import defaultdict
from typing import List

MOD = 1_000_000_007


def count_special_triplets(nums: List[int]) -> int:
    """
    Count triplets (i, j, k) with i < j < k and nums[i] = nums[j]*2 = nums[k].
    Uses prefix/suffix frequency maps to tally matches in O(n).
    """
    suffix = defaultdict(int)
    for v in nums:
        suffix[v] += 1

    prefix = defaultdict(int)
    ans = 0

    for v in nums:
        suffix[v] -= 1
        target = v * 2
        ans = (ans + prefix[target] * suffix[target]) % MOD
        prefix[v] += 1

    return ans


class Solution:
    # LeetCode-style method name for compatibility.
    def specialTriplets(self, nums: List[int]) -> int:
        dic = {}
        dic2 = {}
        res = 0
        for num in nums:
            if num % 2 == 0 and num // 2 in dic2:
                res += dic2[num // 2]
                res %= 10**9 + 7
            if num * 2 in dic:
                if num in dic2:
                    dic2[num] += dic[num * 2]
                else:
                    dic2[num] = dic[num * 2]
            if num in dic:
                dic[num] += 1
            else:
                dic[num] = 1
        return res

    # Convenience wrapper that delegates to the shared implementation above.
    def count_special_triplets(self, nums: List[int]) -> int:
        return count_special_triplets(nums)


if __name__ == "__main__":
    solution = Solution()
    samples = [
        [6, 3, 6],
        [0, 1, 0, 0],
        [8, 4, 2, 8, 4],
    ]
    for nums in samples:
        print(f"nums={nums} -> {solution.count_special_triplets(nums)}")
        print(f"nums={nums} -> {solution.specialTriplets(nums)}")
