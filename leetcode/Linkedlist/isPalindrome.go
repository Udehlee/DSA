package linkedlist

type ListNode struct {
	val  int
	next *ListNode
}

// Find the middle of the linked list using slow and fast pointers
// reverselist
// comparelist
func isPalindrome(head *ListNode) bool {
	if head == nil && head.next == nil {
		return true
	}

	slowPtr := head
	fastPtr := head

	for fastPtr != nil && fastPtr.next != nil {
		slowPtr = slowPtr.next
		fastPtr = fastPtr.next.next
	}

	slowPtr = reverseList(slowPtr)

	return compareLists(head, slowPtr)

}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		next := head.next
		head.next = prev
		prev = head
		head = next
	}
	return prev
}

func compareLists(head1, head2 *ListNode) bool {
	for head2 != nil {
		if head1.val != head2.val {
			return false
		}
		head1 = head1.next
		head2 = head2.next
	}
	return true
}
