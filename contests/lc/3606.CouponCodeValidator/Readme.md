# [3606. Coupon Code Validator](https://leetcode.com/problems/coupon-code-validator/description)

You are given three arrays of length `n` that describe `n` coupons: `code`, `businessLine`, and `isActive`. For each coupon `i`:
- `code[i]`: coupon identifier string
- `businessLine[i]`: business category string
- `isActive[i]`: boolean indicating whether the coupon is active

A coupon is **valid** when all conditions hold:
- `code[i]` is non-empty and uses only alphanumerics or underscores (`a-z`, `A-Z`, `0-9`, `_`).
- `businessLine[i]` is one of: `electronics`, `grocery`, `pharmacy`, `restaurant`.
- `isActive[i]` is `true`.

Return the codes of all valid coupons, sorted by `businessLine` in the order:
`electronics` → `grocery` → `pharmacy` → `restaurant`, and lexicographically by `code` within each category.

## Examples

**Example 1**
```text
Input:
  code = ["SAVE20", "", "PHARMA5", "SAVE@20"]
  businessLine = ["restaurant", "grocery", "pharmacy", "restaurant"]
  isActive = [true, true, true, true]

Output: ["PHARMA5", "SAVE20"]
Explanation:
- First coupon is valid.
- Second coupon has empty code (invalid).
- Third coupon is valid.
- Fourth coupon has special character @ (invalid).
```

**Example 2**
```text
Input:
  code = ["GROCERY15", "ELECTRONICS_50", "DISCOUNT10"]
  businessLine = ["grocery", "electronics", "invalid"]
  isActive = [false, true, true]

Output: ["ELECTRONICS_50"]
Explanation:
- First coupon is inactive (invalid).
- Second coupon is valid.
- Third coupon has invalid business line (invalid).
```

## Constraints
- `n == code.length == businessLine.length == isActive.length`
- `1 <= n <= 100`
- `0 <= code[i].length, businessLine[i].length <= 100`
- `code[i]` and `businessLine[i]` use printable ASCII characters.
- `isActive[i]` is either `true` or `false`.
