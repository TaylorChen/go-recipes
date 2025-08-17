package algorithm

import "fmt"

type Tree struct {
	Left  *Tree
	Right *Tree
	Data  int
}

func NewNode(data int) (bree *Tree) {
	return &Tree{nil, nil, data}
}

func Insert(node *Tree, data int) (bree *Tree) {
	if node == nil {
		return NewNode(data)
	}

	if data < node.Data {
		node.Left = Insert(node.Left, data)
	} else if data > node.Data {
		node.Right = Insert(node.Right, data)
	}
	return node
}

func InOrder(node *Tree) {
	if node != nil {
		InOrder(node.Left)
		fmt.Println(node.Data)
		InOrder(node.Right)
	}

}

func minValueNode(node *Tree) (tree *Tree) {
	current := node

	for current.Left != nil {
		current = current.Left
	}
	return current
}

// func main() {}
