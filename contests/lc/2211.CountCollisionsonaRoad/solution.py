class Solution:
    def countCollisions(self, directions: str) -> int:
        """
        Trim outward-moving cars (leading L's, trailing R's); all remaining
        moving cars will collide and become stationary.
        """
        n = len(directions)

        left = 0
        while left < n and directions[left] == "L":
            left += 1

        right = n - 1
        while right >= left and directions[right] == "R":
            right -= 1

        collisions = 0
        for i in range(left, right + 1):
            if directions[i] != "S":
                collisions += 1

        return collisions
