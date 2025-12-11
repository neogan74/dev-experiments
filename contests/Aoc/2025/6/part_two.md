# Part Two

The cephalopods realize they forgot to explain how to *read* their math. Problems are still separated by a full column of spaces and use the operator in the bottom row, but numbers are written differently.

## Updated rules
- Math is written right-to-left, one column per number.
- Within a column, digits run from top (most significant) to bottom (least significant).
- Read columns from right to left, build the numbers, then apply the block's operator (`+` or `*`).
- The worksheet answer is the sum of each block's result.

## Example
```
123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  
```

Reading by columns, right-to-left:
- Rightmost problem: `4 + 431 + 623 = 1058`
- Second from right: `175 * 581 * 32 = 3253600`
- Third from right: `8 + 248 + 369 = 625`
- Leftmost problem: `356 * 24 * 1 = 8544`

Grand total: `1058 + 3253600 + 625 + 8544 = 3263827`.

Solve the worksheet again using these rules. What is the grand total?
