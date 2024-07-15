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
}
