package main

import "fmt"

type HashTable struct {
	arr  []*Node
	size int
}

type Node struct {
	value int
	next  *Node
}

func (h *HashTable) Hash(key int) int {
	return key % h.size
}

func (h *HashTable) Init(size int) *HashTable {
	h.size = size
	h.arr = make([]*Node, h.size)
	for i := 0; i < h.size; i++ {
		h.arr[i] = nil
	}
	return h
}

func (h *HashTable) Add(value int) {
	index := h.Hash(value)
	curr := new(Node)
	curr.value = value
	curr.next = h.arr[index]
	h.arr[index] = curr
}

func (h *HashTable) Search(value int) (int, bool) {
	index := h.Hash(value)
	head := h.arr[index]
	for head != nil {
		if head.value == value {
			return head.value, true
		}
		head = head.next
	}
	return 0, false
}

func (h *HashTable) Print() {
	for i := 0; i < len(h.arr); i++ {
		node := h.arr[i]
		for node != nil {
			fmt.Println(node.value)
			node = node.next
		}
	}
}

func main() {
	h := new(HashTable)
	h.Init(30)
	h.Add(34)
	h.Add(54)
	h.Print()
	result, found := h.Search(34)
	fmt.Printf("found %v, result: %v\n", found, result)

}
