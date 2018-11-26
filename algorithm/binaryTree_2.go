package main

import (
	"fmt"
	"io"
	"os"
)

type BinaryNode struct {
	left  *BinaryNode
	right *BinaryNode
	data  int64
}

type BinaryTree struct {
	root *BinaryNode
}

func (t *BinaryTree) insert(data int64) *BinaryTree {

	if t.root == nil {
		t.root = &BinaryNode{data: data, left: nil, right: nil}
	} else {
		t.root.insert(data)
	}
	return t
}

func (n *BinaryNode) insert(data int64) {

	if n == nil {
		return
	} else if data <= n.data {
		if n.left == nil {
			n.left = &BinaryNode{data: data, left: nil, right: nil}
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &BinaryNode{data: data, left: nil, right: nil}
		} else {
			n.right.insert(data)
		}
	}
}

func minNode(current *BinaryNode) (node *BinaryNode) {
	for current.left != nil {
		current = current.left
	}
	return current
}

func (n *BinaryTree) search(data int64) (node *BinaryNode) {
	current := n.root
	for current.data != data {
		if current != nil {
			if current.data > data {
				current = current.left
			} else {
				current = current.right
			}
			if current == nil {
				return nil
			}
		}
	}
	return current
}

func deleteNode(root *BinaryNode, data int64) (node *BinaryNode) {
	if root == nil {
		return root
	}

	if data < root.data {
		root.left = deleteNode(root.left, data)
	} else if data > root.data {
		root.right = deleteNode(root.right, data)
	} else {
		if root.left == nil {
			tmp := root.right
			root = nil
			return tmp
		} else if root.right == nil {
			tmp := root.left
			root = nil
			return tmp
		}

		tmp := minNode(root.right)
		root.data = tmp.data
		root.right = deleteNode(root.right, tmp.data)
	}
	return root
}

func preOrder(w io.Writer, node *BinaryNode, ns int, ch rune) {
	if node == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}
	fmt.Fprintf(w, "%c:%v\n", ch, node.data)
	preOrder(w, node.left, ns+2, 'L')
	preOrder(w, node.right, ns+2, 'R')
}

func main() {
	tree := &BinaryTree{}
	tree.insert(100).
		insert(-20).
		insert(-50).
		insert(-15).
		insert(-60).
		insert(50).
		insert(60).
		insert(55).
		insert(85).
		insert(15).
		insert(5).
		insert(-10)
	preOrder(os.Stdout, tree.root, 0, 'M')

	min := minNode(tree.root)
	fmt.Println(min.data)

	searchNode := tree.search(1000)
	if searchNode != nil {
		fmt.Println(searchNode.data)
	} else {
		fmt.Println(searchNode)
	}
	root := deleteNode(tree.root, 15)
	preOrder(os.Stdout, root, 0, 'M')
}
