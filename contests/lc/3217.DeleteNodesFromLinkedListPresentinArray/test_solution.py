import pytest

from solution import build_linked_list, list_to_pylist, modified_list


@pytest.mark.parametrize(
    ("nums", "values", "expected"),
    [
        ([1, 4, 3], [2, 1, 4, 2, 3, 7], [2, 2, 7]),
        ([5, 6], [5, 6, 7, 8], [7, 8]),
        ([9, 10], [1, 2, 3], [1, 2, 3]),
        ([1, 2, 3], [1, 2, 3, 2], []),
        ([2, 2, 3], [1, 2, 3, 4], [1, 4]),
    ],
)
def test_modified_list(nums, values, expected):
    head = build_linked_list(values)
    result = modified_list(nums, head)
    assert list_to_pylist(result) == expected


def test_modified_list_none_head():
    assert modified_list([1, 2], None) is None
