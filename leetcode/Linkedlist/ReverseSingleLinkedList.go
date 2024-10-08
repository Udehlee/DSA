package main

import "fmt"

type node struct {
	value int
	next  *node
}

//iterative
func iterative(head *node) *node {
	var prev, next *node
	curr := head

	for curr != nil {
		next = curr.next
		curr.next = prev
		prev = curr
		curr = next

	}
	return prev
}

//checks if the list is empty or you've reached the last node
func recursive(node *node) *node {
	if node == nil || node.next == nil {
		return node

	}

	reverseList := recursive(node.next)
	node.next.next = node
	node.next = nil

	return reverseList

}

func Print(node *node) {
	toPrint := node

	for toPrint != nil {
		fmt.Print(toPrint.value, " ")
		toPrint = toPrint.next
	}
	fmt.Println("")
}

func main() {
	headd := &node{1, &node{2, nil}}
	fmt.Println()
	print(headd)

	iterativeHead := iterative(headd)
	fmt.Println("iterative")
	print(iterativeHead)

	recursiveHead := recursive(headd)
	fmt.Println("recursive")
	print(recursiveHead)
}
