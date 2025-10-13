from typing import List

class Solution:
    def removeAnagrams(self, words: List[str]) -> List[str]:
        res = []
        prev_key = None
        for w in words:
            key = "".join(sorted(w))
            if key != prev_key:
                res.append(w)
                prev_key = key
        return res