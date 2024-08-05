package main

import "fmt"

//two references in each node in doubly linked list
type DoublyLinkedList struct {
	head  *Node
	tail  *Node
	count int
}
type Node struct {
	prev  *Node
	value int
	next  *Node
}

//Prepend adds a new head to the list
func (l *DoublyLinkedList) Prepend(v int) {
	newAdd := &Node{
		prev:  nil,
		value: v,
		next:  nil,
	}

	if l.count == 0 {
		l.head = newAdd
		l.tail = newAdd
	} else {
		l.head.prev = newAdd
		newAdd.next = l.head
		l.head = newAdd
	}
	l.count++

}

//PrintLinkedList prints the entire list
func (l *DoublyLinkedList) PrintLinkedList() {
	toPrint := l.head

	for toPrint != nil {
		fmt.Print(toPrint.value, " ")
		toPrint = toPrint.next
	}
	fmt.Println("")
}

func main() {

	list := &DoublyLinkedList{}

	list.Prepend(6)
	list.Prepend(1)
	list.Prepend(2)
	list.Prepend(3)
	list.PrintLinkedList()

}
