package linkedlist

type ListNode struct {
	value int
	next  *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	slowPtr := headA
	fastPtr := headB

	for slowPtr != fastPtr {
		if slowPtr == nil {
			slowPtr = headB
		} else {
			slowPtr = slowPtr.next
		}

		if fastPtr == nil {
			fastPtr = headA
		} else {
			fastPtr = fastPtr.next
		}

	}
	return slowPtr
}
