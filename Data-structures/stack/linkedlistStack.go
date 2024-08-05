package main

type Stack struct {
	head  *Node
	count int
}

type Node struct {
	value int
	next  *Node
}

func (s *Stack) Push(value int) {
	s.head = &Node{value, s.head}
	s.count++
}

func (s *Stack) Pop() int {
	toRemove := s.head.value
	s.head = s.head.next
	return toRemove
}
