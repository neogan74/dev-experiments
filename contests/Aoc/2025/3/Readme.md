# Day 3: Lobby

You descend a short staircase into a vast lobby, clear security, and find every elevator offline. An Elf at a control panel blames an electrical surge; while she fixes the elevators, you can try to power the escalator using nearby batteries.

Each battery is labeled with a joltage rating `1-9`. Batteries are arranged in banks, where each line of digits represents one bank. Within a bank you must turn on exactly two batteries; the bank's joltage is the two-digit number formed by those chosen batteries **in their existing order**.

## Example input

```
987654321111111
811111111111119
234234234234278
818181911112111
```

## Finding the bank maximum

- `987654321111111` → best pick is the first two batteries for `98`.
- `811111111111119` → pick `8` and `9` for `89`.
- `234234234234278` → pick the last two batteries for `78`.
- `818181911112111` → pick `9` and `2` for `92`.

Summing those maximums gives a total output joltage of `98 + 89 + 78 + 92 = 357`.

## Task

For your puzzle input, compute the largest possible joltage for each bank and report the total output joltage.
