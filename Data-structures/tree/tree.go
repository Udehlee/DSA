package main

import "fmt"

type Node struct {
	key         int
	left, right *Node
}

//insert adds a key to the tree
//Key goes to left if its smaller or right if its bigger
func (n *Node) insert(k int) {
	if n.key < k {
		if n.right == nil {
			n.right = &Node{key: k}
		} else {
			n.right.insert(k)
		}
	} else if n.key > k {
		if n.left == nil {
			n.left = &Node{key: k}
		} else {
			n.left.insert(k)
		}

	}

}

func (n *Node) Search(k int) bool {
	if n == nil {
		return false
	}
	if n.key < k {
		//move right
		return n.right.Search(k)
	} else if n.key > k {
		//move left
		return n.left.Search(k)
	}
	return true
}
func main() {
	tree := &Node{key: 400}
	tree.insert(200)
	tree.insert(100)
	fmt.Print(tree)

	fmt.Print(tree.Search(100))
}
