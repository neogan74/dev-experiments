# Explanation: Count Collisions on a Road

## Key Insight
Only inward-moving cars can collide. Leading `L` cars move left forever without meeting anyone; trailing `R` cars move right forever. After removing those, every remaining moving car (`L` or `R`) is guaranteed to collide and become stationary. Stationary cars (`S`) inside the trimmed window do not add collisions themselves.

## Algorithm
1. Convert `directions` to a byte/slice/string and find the first index `left` that is not `L`.
2. Find the last index `right` that is not `R`, ensuring `right >= left`.
3. From `left` to `right`, count characters that are not `S`; each such car contributes exactly one collision.
4. Return that count.

## Correctness Proof
We prove the algorithm returns the true collision count.

- Leading `L` cars (before `left`) always move left with no cars to their left, so they never collide. Removing them does not change collisions.
- Trailing `R` cars (after `right`) always move right with no cars to their right, so they never collide. Removing them does not change collisions.
- In the remaining window `[left, right]`, there is at least one non-`L` on the left boundary and one non-`R` on the right boundary, so no car can escape outward.
- Any moving car (`L` or `R`) within this window must eventually meet a stationary car or an opposing moving car, after which it becomes stationary and contributes exactly one collision to the total.
- Stationary cars (`S`) do not initiate collisions but can be collided into; counting only non-`S` cars in the window therefore equals the total collisions that occur.

Thus, the algorithm's count matches the actual number of collisions.

## Complexity Analysis
- Time: O(n) to scan from both ends and count.
- Space: O(1) extra space.

## Reference Implementation
- Go: `solution.go` (`countCollisions`)
- Python: `solution.py` (`Solution.countCollisions`)
