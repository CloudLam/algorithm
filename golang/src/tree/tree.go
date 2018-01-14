// Copyright 2018 Cloud Lam. All rights reserved
// Package tree implements a binary tree

package tree

// Node is a struct of tree node
type Node struct {
	value int
	left  *Node
	right *Node
}

// BinaryTree is a struct of tree
type BinaryTree struct {
	root Node
}
