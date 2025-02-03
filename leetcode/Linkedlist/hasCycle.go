package linkedlist

type ListNode struct {
	value int
	next  *ListNode
}

func hasCycle(head *ListNode) bool {
	slowPtr := head
	fastPtr := head

	for fastPtr != nil && fastPtr.next != nil {
		slowPtr = slowPtr.next
		fastPtr = fastPtr.next.next

		if slowPtr == fastPtr {
			return true
		}
	}
	return false

}
