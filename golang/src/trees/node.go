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

func (node *Node) remove(value int) {
	if node == nil {
		return
	}
	if node.Value == value {
		if node.Left == nil && node.Right == nil {
			*node = Node{-1, nil, nil}
			return
		}
		if node.Left == nil {
			*node = Node{node.Right.Value, node.Right.Left, node.Right.Right}
			return
		}
		if node.Right == nil {
			*node = Node{node.Left.Value, node.Left.Left, node.Left.Right}
			return
		}
		temp := node.Right
		for temp.Left != nil {
			temp = temp.Left
		}
		*node = Node{temp.Value, node.Left, node.Right}
		node.Right.remove(temp.Value)
		return
	}
	if node.Value > value {
		node.Left.remove(value)
		if node.Left.Value == -1 {
			node.Left = nil
		}
	}
	if node.Value < value {
		node.Right.remove(value)
		if node.Right.Value == -1 {
			node.Right = nil
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
