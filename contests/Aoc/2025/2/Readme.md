# Day 2: Gift Shop

You reach the gift shop, the elevator's only other stop. A cheerful sign thanks you for visiting the North Pole. You can reach the lobby from here, and from there the rest of the base.

An Elf clerk asks for help: a younger Elf added a bunch of invalid product IDs to their database. Most ranges are already checked; you just need to handle the remaining ranges (your puzzle input).

## Input format

- Ranges are separated by commas.
- Each range is `start-end` (inclusive).

Example ranges (wrapped for legibility):

```
11-22,95-115,998-1012,1188511880-1188511890,222220-222224,
1698522-1698528,446443-446449,38593856-38593862,565653-565659,
824824821-824824827,2121212118-2121212124
```

## Invalid ID rule

An invalid ID consists of a sequence of digits repeated twice (no leading zeroes). Examples: `55` (`5` twice), `6464` (`64` twice), `123123` (`123` twice). `0101` never appears because IDs have no leading zeroes.

## Worked example

- `11-22` contains two invalid IDs: `11`, `22`.
- `95-115` contains one invalid ID: `99`.
- `998-1012` contains one invalid ID: `1010`.
- `1188511880-1188511890` contains one invalid ID: `1188511885`.
- `222220-222224` contains one invalid ID: `222222`.
- `1698522-1698528` contains no invalid IDs.
- `446443-446449` contains one invalid ID: `446446`.
- `38593856-38593862` contains one invalid ID: `38593859`.
- The remaining ranges contain no invalid IDs.

The sum of all invalid IDs in this example is `1227775554`.

## Task

Find all invalid IDs in your input ranges and report their sum.

-----
