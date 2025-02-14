package linkedlist

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	slowPtr := headA
	fastPtr := headB

	for slowPtr != fastPtr {
		if slowPtr == nil {
			slowPtr = headB
		} else {
			slowPtr = slowPtr.Next
		}

		if fastPtr == nil {
			fastPtr = headA
		} else {
			fastPtr = fastPtr.Next
		}

	}
	return slowPtr
}
