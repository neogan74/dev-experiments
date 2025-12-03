# Part Two

The clerk finds more invalid IDs. Maybe the young Elf used other patterns.

Now an ID is invalid if it consists of a digit sequence repeated **at least twice** (no leading zeroes). Examples: `12341234` (`1234` ×2), `123123123` (`123` ×3), `1212121212` (`12` ×5), `1111111` (`1` ×7).

Using the same example ranges:

- `11-22` still has two invalid IDs: `11`, `22`.
- `95-115` now has two invalid IDs: `99`, `111`.
- `998-1012` now has two invalid IDs: `999`, `1010`.
- `1188511880-1188511890` still has one invalid ID: `1188511885`.
- `222220-222224` still has one invalid ID: `222222`.
- `1698522-1698528` still contains no invalid IDs.
- `446443-446449` still has one invalid ID: `446446`.
- `38593856-38593862` still has one invalid ID: `38593859`.
- `565653-565659` now has one invalid ID: `565656`.
- `824824821-824824827` now has one invalid ID: `824824824`.
- `2121212118-2121212124` now has one invalid ID: `2121212121`.

The sum of invalid IDs in this example is `4174379265`.

What is the sum for your input under the new rule?
