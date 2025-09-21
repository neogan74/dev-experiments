import pytest

from add_two_numbers import ListNode, Solution


def build_linked_list(values):
    head = None
    current = None
    for value in values:
        node = ListNode(value)
        if head is None:
            head = node
        else:
            current.next = node
        current = node
    return head


def linked_list_to_list(node):
    values = []
    while node:
        values.append(node.val)
        node = node.next
    return values or [0]


def test_add_two_numbers_example():
    l1 = build_linked_list([2, 4, 3])
    l2 = build_linked_list([5, 6, 4])

    result = Solution().addTwoNumbers(l1, l2)

    assert linked_list_to_list(result) == [7, 0, 8]


def test_add_two_numbers_with_carry_chain():
    l1 = build_linked_list([9, 9, 9, 9])
    l2 = build_linked_list([1])

    result = Solution().addTwoNumbers(l1, l2)

    assert linked_list_to_list(result) == [0, 0, 0, 0, 1]


def test_add_two_numbers_when_one_number_is_zero():
    l1 = build_linked_list([0])
    l2 = build_linked_list([7, 3])

    result = Solution().addTwoNumbers(l1, l2)

    assert linked_list_to_list(result) == [7, 3]
