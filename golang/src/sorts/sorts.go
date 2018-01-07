// Copyright 2018 Cloud Lam. All rights reserved
// Package sorts implements some sort methods

package sorts

// BubbleSort sorts array with Bubble Sort method
func BubbleSort(array []int) {
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

// InsertSort sorts array with Insert Sort method
func InsertSort(array []int) {
	var temp int
	for i := 1; i < len(array); i++ {
		if array[i] < array[i-1] {
			temp = array[i]
			for j := i - 1; j >= 0 && array[j] > temp; j-- {
				array[j+1] = array[j]
				array[j] = temp
			}
		}
	}
}

// ShellSort sorts array with Shell's Sort Method
func ShellSort(array []int) {
	var temp int
	for gap := len(array) / 2; gap > 0; gap /= 2 {
		for i := gap; i < len(array); i++ {
			if array[i] < array[i-gap] {
				for j := i - gap; j >= 0 && array[j] > array[j+gap]; j -= gap {
					temp = array[j]
					array[j] = array[j+gap]
					array[j+gap] = temp
				}
			}
		}
	}
}
