// Copyright 2018 Cloud Lam. All rights reserved
// Package tree implements binary trees

package trees

// Node implementation
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func (node *Node) insert(value int) {
	if node == nil {
		node = &Node{value, nil, nil}
	}
	if node.Value > value {
		if node.Left == nil {
			node.Left = &Node{value, nil, nil}
		} else {
			node.Left.insert(value)
		}
	} else {
		if node.Right == nil {
			node.Right = &Node{value, nil, nil}
		} else {
			node.Right.insert(value)
		}
	}
}

func (node *Node) preOrder(result *[]int) {
	if node != nil {
		*result = append(*result, node.Value)
	}
	if node.Left != nil {
		node.Left.preOrder(result)
	}
	if node.Right != nil {
		node.Right.preOrder(result)
	}
}

func (node *Node) inOrder(result *[]int) {
	if node.Left != nil {
		node.Left.inOrder(result)
	}
	if node != nil {
		*result = append(*result, node.Value)
	}
	if node.Right != nil {
		node.Right.inOrder(result)
	}
}

func (node *Node) postOrder(result *[]int) {
	if node.Left != nil {
		node.Left.postOrder(result)
	}
	if node.Right != nil {
		node.Right.postOrder(result)
	}
	if node != nil {
		*result = append(*result, node.Value)
	}
}
