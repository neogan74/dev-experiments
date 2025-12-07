# Day 5: Cafeteria

As the forklifts break through the wall, the Elves are delighted to discover that there was a cafeteria on the other side after all. The kitchen is scrambling to decorate before the dining hall opens, and their brand-new inventory system is getting in the way.

## Problem
The database works with ingredient IDs. It lists ranges of fresh ingredient IDs, then (after a blank line) the available ingredient IDs to check. An ingredient ID is fresh if it falls within any range.

## Input format
- One or more fresh ID ranges, written as `start-end` (inclusive). Ranges may overlap.
- A blank line separator.
- A list of available ingredient IDs, one per line.

## Example
```
3-5
10-14
16-20
12-18

1
5
8
11
17
32
```

Walkthrough:
- 1: spoiled (not in any range)
- 5: fresh (in 3-5)
- 8: spoiled
- 11: fresh (in 10-14)
- 17: fresh (in 16-20 and 12-18)
- 32: spoiled

In this example, 3 of the available ingredient IDs are fresh.

## Task
Process the database file from the new inventory management system and count how many of the available ingredient IDs are fresh.
