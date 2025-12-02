# 3217. Delete Nodes From Linked List Present in Array

## Problem Statement

You are given:
- An integer array `nums`.
- The head of a singly linked list `head`. Each node of the list has an integer value.

Delete every node from the linked list whose value appears in `nums`, and return a reference to the head of the updated list. If every node is removed, return `null`.

The linked list structure is defined as:

```text
class ListNode {
    int val;
    ListNode next;
}
```

### Constraints

- `1 <= nums.length <= 10^5`
- `nums` may contain duplicate values.
- The linked list contains at least one node.
- `-10^5 <= nums[i], head.val <= 10^5`
- The total number of nodes in the linked list is at most `10^5`.

## Example

```text
Input: nums = [1, 4, 3], head = 2 -> 1 -> 4 -> 2 -> 3 -> 7
Output: 2 -> 2 -> 7
Explanation: Nodes with values 1, 4, and 3 are removed from the list.
```

## Solution Overview

1. Insert all values from `nums` into a hash set for `O(1)` lookups.
2. Traverse the linked list while tracking the previous node.
3. Whenever the current node's value exists in the hash set, unlink it by updating the previous node's `next` pointer. Handle deletions at the head by advancing `head`.
4. Continue until the end of the list, then return the new head.

### Complexity

- **Time Complexity:** `O(n + m)` where `n` is the number of values in `nums`, and `m` is the number of nodes in the linked list.
- **Space Complexity:** `O(n)` for storing the values of `nums` in a hash set.

