import importlib.util
import pathlib
import unittest

MODULE_PATH = pathlib.Path(__file__).with_name("count-elevemts-with-maximum-freq.py")


def load_solution():
    """Dynamically load the solution module so the hyphenated filename works under tests."""
    spec = importlib.util.spec_from_file_location("count_elevemts_with_maximum_freq", MODULE_PATH)
    module = importlib.util.module_from_spec(spec)
    loader = spec.loader
    if loader is None:
        raise AssertionError("Unable to load solution module: missing loader")
    try:
        loader.exec_module(module)
    except SyntaxError as exc:  # pragma: no cover - provides clearer failure
        raise AssertionError("Solution module contains a syntax error") from exc
    return module.Solution()


class MaxFrequencyElementsTests(unittest.TestCase):
    def setUp(self):
        self.solution = load_solution()

    def test_mixed_frequencies(self):
        nums = [1, 2, 2, 3, 1, 4]
        self.assertEqual(self.solution.maxFrequencyElements(nums), 4)

    def test_all_elements_equal(self):
        nums = [7, 7, 7, 7]
        self.assertEqual(self.solution.maxFrequencyElements(nums), 4)

    def test_multiple_high_frequency_groups(self):
        nums = [1, 2, 2, 3, 3, 4, 4]
        self.assertEqual(self.solution.maxFrequencyElements(nums), 6)


if __name__ == "__main__":
    unittest.main()
