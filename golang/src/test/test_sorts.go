// Copyright 2018 Cloud Lam. All rights reserved

package main

import (
	"fmt"
	"math/rand"
	"sorts"
	"time"
)

// SIZE is the array length
var SIZE = 10240

// SortTimes prints all sort method time cost
func SortTimes() {
	fmt.Println("The length of array: ", SIZE)
	SortTime("Bubble")
	SortTime("Insert")
	SortTime("Shell")
	SortTime("Select")
	SortTime("Heap")
	SortTime("Quick")
	SortTime("Merge")
}

// SortTime prints sort method time cost
func SortTime(method string) {
	// Build array
	array := []int{}
	for index := 0; index < SIZE; index++ {
		array = append(array, rand.Intn(100000))
	}
	// Begin time
	begin := time.Now()
	switch method {
	case "Bubble":
		sorts.BubbleSort(array)
	case "Insert":
		sorts.InsertSort(array)
	case "Shell":
		sorts.ShellSort(array)
	case "Select":
		sorts.SelectSort(array)
	case "Heap":
		sorts.HeapSort(array)
	case "Quick":
		sorts.QuickSort(array, 0, SIZE-1)
	case "Merge":
		sorts.MergeSort(array)
	default:
	}
	// Cost time
	cost := time.Since(begin)
	fmt.Println(method+" Sort Cost: ", cost)
}

// BubbleSortTest tests BubbleSort
func BubbleSortTest() {
	array := []int{89, 64, 18, 37, 95, 73, 25, 56, 43, 100}
	fmt.Println("Source: ", array)
	fmt.Println("Result: ", sorts.BubbleSort(array))
}

// InsertSortTest tests InsertSort
func InsertSortTest() {
	array := []int{89, 64, 18, 37, 95, 73, 25, 56, 43, 100}
	fmt.Println("Source: ", array)
	fmt.Println("Result: ", sorts.InsertSort(array))
}

// ShellSortTest tests ShellSort
func ShellSortTest() {
	array := []int{89, 64, 18, 37, 95, 73, 25, 56, 43, 100}
	fmt.Println("Source: ", array)
	fmt.Println("Result: ", sorts.ShellSort(array))
}

// SelectSortTest tests SelectSort
func SelectSortTest() {
	array := []int{89, 64, 18, 37, 95, 73, 25, 56, 43, 100}
	fmt.Println("Source: ", array)
	fmt.Println("Result: ", sorts.SelectSort(array))
}

// HeapSortTest tests HeapSort
func HeapSortTest() {
	array := []int{89, 64, 18, 37, 95, 73, 25, 56, 43, 100}
	fmt.Println("Source: ", array)
	fmt.Println("Result: ", sorts.HeapSort(array))
}

// QuickSortTest tests QuickSort
func QuickSortTest() {
	array := []int{89, 64, 18, 37, 95, 73, 25, 56, 43, 100}
	fmt.Println("Source: ", array)
	fmt.Println("Result: ", sorts.QuickSort(array, 0, len(array)-1))
}

// MergeSortTest tests MergeSort
func MergeSortTest() {
	array := []int{89, 64, 18, 37, 95, 73, 25, 56, 43, 100}
	fmt.Println("Source: ", array)
	fmt.Println("Result: ", sorts.MergeSort(array))
}
