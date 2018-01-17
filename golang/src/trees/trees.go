// Copyright 2018 Cloud Lam. All rights reserved
// Package tree implements binary trees

package trees

// BinarySearchTree implementation
// BinarySearchTree is a struct of tree
type BinarySearchTree struct {
	Root *Node
}

// Insert adds value into the Binary-Search-Tree
func (bst *BinarySearchTree) Insert(value int) {
	bst.Root.insert(value)
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
	if data == nil {
		return &BinarySearchTree{&Node{}}
	}
	bst := &BinarySearchTree{&Node{data[0], nil, nil}}
	for i := 1; i < len(data); i++ {
		bst.Insert(data[i])
	}
	return bst
}
