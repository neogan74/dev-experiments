from typing import Optional


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

class Solution:
    def addTwoNumbers(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
        dummy = ListNode(0)
        curr = dummy
        carry = 0

        while l1 or l2 or carry:
            val1 = l1.val if l1 else 0
            val2 = l2.val if l2 else 0

            total = val1 + val2 + carry
            carry = total // 10
            curr.next = ListNode(total % 10)
            curr = curr.next

            if l1:
                l1 = l1.next
            if l2:
                l2 = l2.next

        return dummy.next
    
def main():
    # Example usage:
    l1 = ListNode(2, ListNode(4, ListNode(3)))  # Represents the number 342
    l2 = ListNode(5, ListNode(6, ListNode(4)))  # Represents the number 465
    result = Solution().addTwoNumbers(l1, l2)
    
    # Print the result
    while result:
        print(result.val, end=" -> ")
        result = result.next
    print("None")

if __name__ == "__main__":
    main()