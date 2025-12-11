# Day 6: Trash Compactor

After an enthusiastic jump into a garbage chute, you land in a magnetically sealed trash compactor. While waiting for a family of cephalopods to open the door, you are asked to help with a math worksheet.

Each column block in the worksheet represents a problem: the top rows contain the numbers, and the bottom row holds the operator (`+` or `*`). Problems sit side by side, separated by at least one full column of spaces.

## Input
- A rectangular grid of ASCII characters.
- The last row contains the operators for each problem.
- Columns that are all spaces separate problems.
- Within a problem block, left/right alignment of the numbers does not matter.

## Output
Print the grand total: evaluate every problem using its operator, then sum all of those results.

## Example
```
123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  
```

Problems and answers:
- `123 * 45 * 6 = 33210`
- `328 + 64 + 98 = 490`
- `51 * 387 * 215 = 4243455`
- `64 + 23 + 314 = 401`

Grand total: `33210 + 490 + 4243455 + 401 = 4277556`.
