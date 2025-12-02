from __future__ import annotations

from dataclasses import dataclass
from typing import Iterable, Optional, Set


@dataclass
class ListNode:
    val: int
    next: Optional["ListNode"] = None


def modified_list(nums: Iterable[int], head: Optional[ListNode]) -> Optional[ListNode]:
    """Remove every node from the list whose value appears in nums."""
    if head is None:
        return None

    blocked: Set[int] = set(nums)

    dummy = ListNode(0, head)
    prev = dummy
    curr = head

    while curr:
        if curr.val in blocked:
            prev.next = curr.next
        else:
            prev = curr
        curr = curr.next

    return dummy.next


def build_linked_list(values: Iterable[int]) -> Optional[ListNode]:
    """Utility to construct a linked list from an iterable of ints."""
    iterator = iter(values)
    try:
        first = next(iterator)
    except StopIteration:
        return None
    head = ListNode(first)
    tail = head
    for value in iterator:
        tail.next = ListNode(value)
        tail = tail.next
    return head


def list_to_pylist(head: Optional[ListNode]) -> list[int]:
    """Utility to convert a linked list back into a Python list."""
    out: list[int] = []
    while head:
        out.append(head.val)
        head = head.next
    return out
