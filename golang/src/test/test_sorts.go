// Copyright 2018 Cloud Lam. All rights reserved

package main

import (
	"fmt"
	"math/rand"
	"sorts"
	"time"
)

// SortTimes prints all sort method time cost
func SortTimes() {
	SortTime("Bubble")
	SortTime("Insert")
	SortTime("Shell")
	SortTime("Select")
}

// SortTime prints sort method time cost
func SortTime(method string) {
	// Build array
	array := []int{}
	for index := 0; index < 10240; index++ {
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
	sorts.BubbleSort(array)
	fmt.Println("Result: ", array)
}

// InsertSortTest tests InsertSort
func InsertSortTest() {
	array := []int{89, 64, 18, 37, 95, 73, 25, 56, 43, 100}
	fmt.Println("Source: ", array)
	sorts.InsertSort(array)
	fmt.Println("Result: ", array)
}

// ShellSortTest tests ShellSort
func ShellSortTest() {
	array := []int{89, 64, 18, 37, 95, 73, 25, 56, 43, 100}
	fmt.Println("Source: ", array)
	sorts.ShellSort(array)
	fmt.Println("Result: ", array)
}

// SelectSortTest tests SelectSort
func SelectSortTest() {
	array := []int{89, 64, 18, 37, 95, 73, 25, 56, 43, 100}
	fmt.Println("Source: ", array)
	sorts.SelectSort(array)
	fmt.Println("Result: ", array)
}
