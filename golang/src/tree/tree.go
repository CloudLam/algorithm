// Copyright 2018 Cloud Lam. All rights reserved
// Package tree implements a binary tree

package tree

// BinarySearchTree is a struct of tree
type BinarySearchTree struct {
	value int
	left  *BinarySearchTree
	right *BinarySearchTree
}

func (bst *BinarySearchTree) search(value int) *BinarySearchTree {
	current := bst
	for current != nil {
		if current.value == value {
			return current
		}
		if current.value > value {
			current = current.left
		} else {
			current = current.right
		}
	}
	return nil
}

func (bst *BinarySearchTree) insert(value int) *BinarySearchTree {
	if bst == nil {
		return &BinarySearchTree{value, nil, nil}
	}
	if value < bst.value {
		bst.left = bst.left.insert(value)
	} else {
		bst.right = bst.right.insert(value)
	}
	return bst
}

// NewBinarySearchTree is the constructor of BinarySearchTree
func NewBinarySearchTree(data []int) *BinarySearchTree {
	if len(data) == 0 {
		return &BinarySearchTree{}
	}
	bst := &BinarySearchTree{data[0], nil, nil}
	for i := 1; i < len(data); i++ {
		bst.insert(data[i])
	}
	return bst
}
