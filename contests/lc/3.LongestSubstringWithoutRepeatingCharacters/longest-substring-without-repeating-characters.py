from typing import Dict


def length_of_longest_substring(s: str) -> int:
    """Return length of longest substring without repeating characters."""
    last_seen: Dict[str, int] = {}
    start = 0
    max_len = 0

    for idx, ch in enumerate(s):
        if ch in last_seen and last_seen[ch] >= start:
            start = last_seen[ch] + 1

        last_seen[ch] = idx
        window_len = idx - start + 1
        if window_len > max_len:
            max_len = window_len

    return max_len


if __name__ == "__main__":
    examples = (
        "abcabcbb",
        "bbbbb",
        "pwwkew",
        "",
    )
    for example in examples:
        print(f"{example!r}: {length_of_longest_substring(example)}")
