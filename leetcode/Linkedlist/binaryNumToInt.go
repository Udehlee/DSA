package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	output := 0

	for head != nil {
		output = (output << 1) | head.Val
		head = head.Next
	}
	return output

}
