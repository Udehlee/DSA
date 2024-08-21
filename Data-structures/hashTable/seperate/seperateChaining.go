package main

import "fmt"

type HashTable struct {
	arr  []*Node
	size int
}

type Node struct {
	key  string
	next *Node
}

func (h *HashTable) Hash(key string) int {
	hash := 0
	for _, v := range key {
		hash = int(v) % h.size
	}
	return hash
}

func (h *HashTable) Init(size int) *HashTable {
	h.size = size
	h.arr = make([]*Node, h.size)
	for i := 0; i < h.size; i++ {
		h.arr[i] = nil
	}
	return h
}

func (h *HashTable) Add(key string) {
	index := h.Hash(key)
	curr := new(Node)
	curr.key = key
	curr.next = h.arr[index]
	h.arr[index] = curr
}

func (h *HashTable) Search(key string) (string, bool) {
	index := h.Hash(key)
	head := h.arr[index]
	for head != nil {
		if head.key == key {
			return head.key, true
		}
		head = head.next
	}
	return "", false
}

func (h *HashTable) Print() {
	for i := 0; i < len(h.arr); i++ {
		node := h.arr[i]
		for node != nil {
			fmt.Println(node.key)
			node = node.next
		}
	}
}

func main() {
	h := &HashTable{}
	h.Init(30)
	h.Add("peter")
	h.Add("ada")
	h.Print()
	result, found := h.Search("peter")
	fmt.Printf("found %v, result: %v\n", found, result)

}
