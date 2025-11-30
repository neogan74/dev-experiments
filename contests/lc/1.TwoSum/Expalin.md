# Two Sum Algorithm Explainer

This solution uses a single-pass hash map to find two indices whose values sum to the target.

## Core Idea
- As we scan `nums`, we know the current value `v` needs a partner `target - v`.
- If that partner has been seen before, we already stored its index and can return the pair immediately.
- If not, we remember the current value and its index so later elements can find it as their partner.

## Step-by-Step
1. Initialize an empty map `value -> index`.
2. For each index `i` and value `v` in `nums`:
   - Compute `complement = target - v`.
   - If `complement` exists in the map, return `[map[complement], i]`.
   - Otherwise store `map[v] = i` and continue.
3. Return an empty list or `nil` if no pair is found (per problem statement this should not happen).

## Correctness
- When we process `v`, all earlier elements are already in the map. If any earlier element can pair with `v`, its value equals `complement` and will be found in the lookup.
- The first time a valid pair is encountered, we return their indices; later elements cannot invalidate that answer because the problem guarantees a single solution.

## Complexity
- Time: `O(n)` because each element is processed once with `O(1)` average hash lookups.
- Space: `O(n)` for the map storing seen values.

## Notes on the Implementations
- Python: `Solution.twoSum` and `twoSum2` implement the logic in a single pass; `twoSum2` uses the assignment expression for compactness.
- Go: `twoSum` mirrors the same approach with a map and short declaration `if` to check the complement in place.

## Example
For `nums = [2, 7, 11, 15], target = 9`:
- `i=0, v=2`: complement `7` not in map → store `{2: 0}`.
- `i=1, v=7`: complement `2` is in map at index `0` → return `[0, 1]`.
