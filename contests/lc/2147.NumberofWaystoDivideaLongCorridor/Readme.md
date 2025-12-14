# [2147. Number of Ways to Divide a Long Corridor](https://leetcode.com/problems/number-of-ways-to-divide-a-long-corridor/description)

Along a library corridor you must place dividers so that each section contains exactly two seats (`S`) and any number of plants (`P`). There is already a divider before index `0` and after index `n-1`; for every gap between characters at most one more divider can be added. Return the number of valid ways modulo `1e9+7`, or `0` if it is impossible.

## Examples

- `corridor = "SSPPSPS"` → `3`
- `corridor = "PPSPSP"` → `1`
- `corridor = "S"` → `0`

See `image-2.png`, `image-1.png`, and `image.png` for visualizations.

## Approach

1. Collect the indices of all seats. If there are `0` seats or an odd number of seats, no valid division exists.
2. Seats are grouped in pairs. The first pair fixes the first section. For every subsequent pair, any divider placed between the last seat of the previous pair and the first seat of the current pair yields a valid partition.
3. If the first seat of a pair is at index `a` and the previous seat (end of the last pair) is at index `b`, there are `(a - b)` possible divider positions (all plants plus the gap immediately after `b`).
4. Multiply these counts for all pairs and take the result modulo `1e9+7`.

### Why multiplying gaps works

- A valid section must contain exactly two seats, so seat indices must be consumed in order: `(S0, S1)` make the first section, `(S2, S3)` make the second, and so on.
- Once `(S0, S1)` are fixed, the divider that ends the first section can be placed in any position from immediately after `S1` up to just before `S2`. Every plant in this stretch plus the gap right after `S1` is a distinct legal placement.
- These choices are independent between pairs, so the total number of layouts is the product of the available divider counts between each consecutive pair of seat-pairs.

## Complexity

- Time: `O(n)` to scan the corridor once.
- Space: `O(k)` to store seat indices (`k` is the seat count), which is `O(n)` in the worst case.

## Code

- Go solution: `solution.go`
- Python solution: `solution.py`
