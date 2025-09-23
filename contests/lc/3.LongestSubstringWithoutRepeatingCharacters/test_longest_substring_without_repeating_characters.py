import importlib.util
import pathlib
import unittest

_MODULE_PATH = pathlib.Path(__file__).with_name("longest-substring-without-repeating-characters.py")
_spec = importlib.util.spec_from_file_location("longest_substring", _MODULE_PATH)
_module = importlib.util.module_from_spec(_spec)
assert _spec.loader is not None
_spec.loader.exec_module(_module)
length_of_longest_substring = _module.length_of_longest_substring


class LengthOfLongestSubstringTests(unittest.TestCase):

    def test_examples(self) -> None:
        self.assertEqual(length_of_longest_substring("abcabcbb"), 3)
        self.assertEqual(length_of_longest_substring("bbbbb"), 1)
        self.assertEqual(length_of_longest_substring("pwwkew"), 3)

    def test_empty_string(self) -> None:
        self.assertEqual(length_of_longest_substring(""), 0)

    def test_all_unique(self) -> None:
        self.assertEqual(length_of_longest_substring("abcdef"), 6)

    def test_repeating_non_letters(self) -> None:
        self.assertEqual(length_of_longest_substring("a!@#a!"), 4)

    def test_middle_repeat(self) -> None:
        self.assertEqual(length_of_longest_substring("abba"), 2)


if __name__ == "__main__":
    unittest.main()
