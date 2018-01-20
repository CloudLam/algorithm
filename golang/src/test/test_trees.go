// Copyright 2018 Cloud Lam. All rights reserved

package main

import (
	"fmt"
	"trees"
)

// BinarySearchTreeTest tests BinarySearchTree
func BinarySearchTreeTest() {
	data := []int{7, 3, 8, 1, 2, 9, 5, 6, 10, 4}
	bst := trees.NewBinarySearchTree(data)
	fmt.Println("Preorder Traversal: ", bst.PreOrder())
	fmt.Println("Inorder Traversal: ", bst.InOrder())
	fmt.Println("Postorder Traversal: ", bst.PostOrder())
	fmt.Println("Minimum: ", bst.GetMin())
	fmt.Println("Maximum: ", bst.GetMax())
	bst.Remove(10)
	fmt.Println(bst.InOrder())
}
