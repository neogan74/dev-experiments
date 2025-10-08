from bisect import bisect_left
from typing import List


def successful_pairs(spells: List[int], potions: List[int], success: int) -> List[int]:
    """Return counts of potions pairing with each spell to reach the success threshold."""
    potions.sort()
    m = len(potions)
    result: List[int] = []

    for spell in spells:
        if spell == 0:
            result.append(0)
            continue

        # minimum potion strength needed; equivalent to math.ceil(success / spell)
        threshold = (success + spell - 1) // spell
        idx = bisect_left(potions, threshold)
        result.append(m - idx)

    return result


class Solution:
    def successfulPairs(self, spells: List[int], potions: List[int], success: int) -> List[int]:
        return successful_pairs(spells, potions, success)
