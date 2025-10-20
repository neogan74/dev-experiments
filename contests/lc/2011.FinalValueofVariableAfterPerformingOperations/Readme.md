# [2011. Final Value of Variable After Performing Operations](https://leetcode.com/problems/final-value-of-variable-after-performing-operations/description)

There is a programming language with only one variable `X` and four operations:

- `++X` and `X++` increment `X` by 1.
- `--X` and `X--` decrement `X` by 1.

`X` starts at 0. Given an array of operations, determine the final value of `X` after applying them.

## Examples

- **Input:** `["--X","X++","X++"]`  
  **Output:** `1`

- **Input:** `["++X","++X","X++"]`  
  **Output:** `3`

- **Input:** `["X++","++X","--X","X--"]`  
  **Output:** `0`

## Constraints

- `1 <= operations.length <= 100`
- `operations[i] ∈ {"++X", "X++", "--X", "X--"}`
