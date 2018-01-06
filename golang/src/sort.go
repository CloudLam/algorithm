// Copyright 2018 Cloud Lam. All rights reserved

package main

import "fmt"

func bubbleSort(array []int) {
	var temp int
	for i := 1; i < len(array); i++ {
		for j := 0; j < len(array)-i; j++ {
			if array[j] > array[j+1] {
				temp = array[j]
				array[j] = array[j+1]
				array[j+1] = temp
			}
		}
	}
}

func main() {
	array := []int{89, 64, 18, 37, 95, 73, 25, 56, 43, 100}
	fmt.Println(array)
	bubbleSort(array)
	fmt.Println(array)
}
