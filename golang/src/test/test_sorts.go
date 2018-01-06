// Copyright 2018 Cloud Lam. All rights reserved

package main

import (
	"fmt"
	"sorts"
)

func BubbleSortTest() {
	array := []int{89, 64, 18, 37, 95, 73, 25, 56, 43, 100}
	fmt.Println(array)
	sorts.BubbleSort(array)
	fmt.Println(array)
}
