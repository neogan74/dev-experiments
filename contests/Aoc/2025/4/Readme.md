## Day 4: Printing Department

You ride the escalator down to the printing department. They are clearly getting ready for Christmas with large rolls of paper everywhere and a massive printer in the corner. You need a path deeper into the North Pole base while the elevators are offline, and the Elves think they can help if their forklifts have some spare time to break through a wall.

If you can optimize their work, they will have time to spare.

### Problem
- Rolls of paper are represented by `@` on a large grid (your puzzle input).
- A roll is accessible to a forklift only if fewer than **four** rolls of paper occupy its eight adjacent positions.
- Goal: determine how many rolls are accessible.

### Example Input
```
..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.
```

### Example Output
Accessible rolls are marked with `x` (13 total):
```
..xx.xx@x.
x@@.@.@.@@
@@@@@.x.@@
@.@@@@..@.
x@.@@@@.@x
.@@@@@@@.@
.@.@.@.@@@
x.@@@.@@@@
.@@@@@@@@.
x.x.@@@.x.
```

### Task
Given the complete diagram of roll locations, count how many rolls of paper can be accessed by a forklift.
