package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteMiddle(head *ListNode) *ListNode {
	slowPtr := head
	fastPtr := head
	var prev *ListNode

	for fastPtr != nil && fastPtr.Next != nil {
		prev = slowPtr
		slowPtr = slowPtr.Next
		fastPtr = fastPtr.Next.Next
	}
	prev.Next = slowPtr.Next
	return head

}
