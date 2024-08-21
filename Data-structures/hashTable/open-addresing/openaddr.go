package main

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
	return h
}

// Linear probing to add a value to the hash table
func (h *HashTable) add(key string) {
	index := h.Hash(key)

	for h.arr[index] != nil {
		index = (index + 1) % h.size
		continue
	}
	h.arr[index] = &Node{key: key}
}
