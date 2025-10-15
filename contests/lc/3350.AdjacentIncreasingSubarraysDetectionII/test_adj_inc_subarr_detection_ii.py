from adj_increaseing_sunarr_detection_ii import Solution
import unittest

def test_example():
    nums = [2,5,7,8,9,2,3,4,3,1]
    assert Solution().maxIncreasingSubarrays(nums) == 3

def test_all_increasing_len5():
    assert Solution().maxIncreasingSubarrays([1,2,3,4,5,1]) == 2

def test_plateaus():
    assert Solution().maxIncreasingSubarrays([1,1,1,1]) == 1

def test_non_increasing():
    assert Solution().maxIncreasingSubarrays([5,4,3]) == 1

def test_single():
    assert Solution().maxIncreasingSubarrays([7]) == 0
