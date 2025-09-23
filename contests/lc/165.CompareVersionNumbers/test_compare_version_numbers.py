import importlib.util
from pathlib import Path

import pytest


_MODULE_PATH = Path(__file__).with_name("compare-version-numbers.py")
_MODULE_SPEC = importlib.util.spec_from_file_location("compare_version_numbers", _MODULE_PATH)
_MODULE = importlib.util.module_from_spec(_MODULE_SPEC)
assert _MODULE_SPEC and _MODULE_SPEC.loader  # ensures loader exists for exec
_MODULE_SPEC.loader.exec_module(_MODULE)
Solution = _MODULE.Solution


@pytest.mark.parametrize(
    "version1, version2, expected",
    [
        ("1.2", "1.10", -1),
        ("1.01", "1.001", 0),
        ("1.0", "1.0.0.0", 0),
        ("1.0.1", "1", 1),
        ("7.5.2.4", "7.5.3", -1),
        ("1.01.001", "1.1.1", 0),
    ],
)
def test_compare_version_numbers(version1, version2, expected):
    assert Solution().compareVersion(version1, version2) == expected
