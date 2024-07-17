package main

import "fmt"

//Linkedlsit - list of items called nodes
//nodes has a value and a reference address pointing to the next node
type LinkedList struct {
	head  *Node
	count int
}

type Node struct {
	value int
	next  *Node
}

//size of the list

func (l *LinkedList) size() int {
	return l.count
}

//Empty?
func (l *LinkedList) isEmpty() bool {
	return l.count == 0
}

//add to the list at the beginning
//varable to hold currentHead
//place the new node to the linkedlist head
//new Node points to old head
func (l *LinkedList) Prepend(v int) {

	currentHead := l.head
	newAdd := &Node{
		value: v,
		next:  l.head,
	}

	l.head = newAdd
	l.head.next = currentHead
	l.count++

}

//PrintLinkedList prints the entire list
func (l *LinkedList) PrintLinkedList() {

	toPrint := l.head

	for toPrint != nil {
		fmt.Print(toPrint.value, " ")
		toPrint = toPrint.next
	}
	fmt.Println("")
}

// add at the end of the list
// add new node if there is no node in the current head

func (l *LinkedList) Append(v int) {
	currentHead := l.head

	newAdd := &Node{
		value: v,
		next:  nil,
	}

	if currentHead == nil {
		currentHead = newAdd
	}

	for currentHead.next != nil {
		currentHead = currentHead.next
	}
	currentHead.next = newAdd
}

//search in a single list can only be done in one direction
//
func (l *LinkedList) isPresent(value int) bool {
	toPrint := l.head

	for toPrint != nil {
		if toPrint.value == value {
			return true
		}

		toPrint = toPrint.next
	}

	return false
}

//delete first element in linked list
func (l *LinkedList) DeleteFirstElement() (int, bool) {
	if l.isEmpty() {
		return 0, false
	}

	value := l.head.value
	l.head = l.head.next
	l.count--
	return value, true

}

//delete from the list any value
func (l *LinkedList) DeleteAny(v int) {
	node := l.head

	for node.next.value != v {
		node = node.next
	}
	node.next = node.next.next
	l.count--
}
func main() {

	list := &LinkedList{}

	//prepend to list
	list.Prepend(1)
	list.Prepend(2)
	list.Prepend(3)
	list.PrintLinkedList()

	//append to list
	list.Append(23)
	list.Append(24)
	list.Append(25)
	list.PrintLinkedList()

	//delete first element
	list.DeleteFirstElement()
	list.PrintLinkedList()

	list.DeleteAny(24)
	list.PrintLinkedList()
}
