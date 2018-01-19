// Copyright 2018 Cloud Lam. All rights reserved
// Package tree implements binary trees

package trees

// BinarySearchTree implementation
// BinarySearchTree is a struct of tree
type BinarySearchTree struct {
	Root *Node
}

// Search finds the node's value equals to input
func (bst *BinarySearchTree) Search(value int) *Node {
	current := bst.Root
	for current != nil {
		if current.Value == value {
			return current
		}
		if current.Value > value {
			current = current.Left
		} else {
			current = current.Right
		}
	}
	return nil
}

// GetMin finds the minimun node of the Binary-Search-Tree
func (bst *BinarySearchTree) GetMin() *Node {
	min := bst.Root
	for {
		if min.Left != nil {
			min = min.Left
		} else {
			return min
		}
	}
}

// GetMax finds the maximun node of the Binary-Search-Tree
func (bst *BinarySearchTree) GetMax() *Node {
	max := bst.Root
	for {
		if max.Right != nil {
			max = max.Right
		} else {
			return max
		}
	}
}

// Insert adds value into the Binary-Search-Tree
func (bst *BinarySearchTree) Insert(value int) {
	bst.Root.insert(value)
}

// Remove deletes value from the Binary-Search-Tree
func (bst *BinarySearchTree) Remove(value int) {
	bst.Root.remove(value)
}

// PreOrder returns Preorder Traversal of the Binary-Search-Tree
func (bst *BinarySearchTree) PreOrder() []int {
	var result []int
	bst.Root.preOrder(&result)
	return result
}

// InOrder returns Inorder Traversal of the Binary-Search-Tree
func (bst *BinarySearchTree) InOrder() []int {
	var result []int
	bst.Root.inOrder(&result)
	return result
}

// PostOrder returns Postorder Traversal of the Binary-Search-Tree
func (bst *BinarySearchTree) PostOrder() []int {
	var result []int
	bst.Root.postOrder(&result)
	return result
}

// NewBinarySearchTree is the constructor of Binary-Search-Tree
func NewBinarySearchTree(data []int) *BinarySearchTree {
	if len(data) == 0 {
		return &BinarySearchTree{&Node{}}
	}
	bst := &BinarySearchTree{&Node{data[0], nil, nil}}
	for i := 1; i < len(data); i++ {
		bst.Insert(data[i])
	}
	return bst
}
