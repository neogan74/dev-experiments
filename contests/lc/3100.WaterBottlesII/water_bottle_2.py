
def maxBottlesDrunk(numBottles: int, numExchange: int) -> int:
    """Return maximum bottles that can be drunk under exchange rules."""
    total_drunk = numBottles

    while numBottles >= numExchange:
        numBottles -= numExchange
        numExchange += 1
        total_drunk += 1
        numBottles += 1

    return total_drunk


class Solution:
    def maxBottlesDrunk(self, numBottles: int, numExchange: int) -> int:
        return maxBottlesDrunk(numBottles, numExchange)
