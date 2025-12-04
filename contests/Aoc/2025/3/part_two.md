# Part Two

The escalator still refuses to move. The Elf hits the big red "joltage limit safety override" button repeatedly while you count confirmations and decorate the lobby.

Now, each bank must power on **exactly twelve** batteries. The bank's joltage is the 12-digit number formed by those chosen batteries **in their original order**.

## Example input

```
987654321111111
811111111111119
234234234234278
818181911112111
```

## Bank maximums with 12 batteries

- `987654321111111` → keep everything except some trailing `1`s to make `987654321111`.
- `811111111111119` → keep everything except some `1`s for `811111111119`.
- `234234234234278` → drop a `2`, a `3`, and another early `2` to make `434234234278`.
- `818181911112111` → skip some leading `1`s to make `888911112111`.

Total output joltage: `987654321111 + 811111111119 + 434234234278 + 888911112111 = 3121910778619`.

## Task

For your puzzle input, choose twelve batteries per bank to maximize each bank's joltage and report the total.
