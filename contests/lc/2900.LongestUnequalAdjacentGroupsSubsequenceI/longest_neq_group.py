from typing import List


class Solution:
    def getLongestSubsequence(self, words: List[str], groups: List[int]) -> List[str]:
        ans: List[str] = []
        prev_group = None
        for word, group in zip(words, groups):
            if prev_group is None or group != prev_group:
                ans.append(word)
                prev_group = group
        return ans
