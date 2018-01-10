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

// SelectSort sorts array with Select Sort Method
func SelectSort(array []int) {
	var max, min, temp int
	for i := 0; i < len(array)/2; i++ {
		max = i
		min = i
		for j := i + 1; j < len(array)-i; j++ {
			if array[j] > array[max] {
				max = j
				continue
			}
			if array[j] < array[min] {
				min = j
			}
		}
		// Minimum
		if min != i {
			temp = array[i]
			array[i] = array[min]
			array[min] = temp
		}
		// Maximum
		if max != i {
			temp = array[len(array)-i-1]
			array[len(array)-i-1] = array[max]
			array[max] = temp
		}
	}
}

// HeapAdjust adjust array
// The array must conform to a heap except array[pos],
// adjust array[pos] to make sure array[i] >= array[2i+1] && array[i] >= array[2i+2].
func HeapAdjust(array []int, pos int, length int) {
	var temp = array[pos]
	for i := pos*2 + 1; i < length; i = i*2 + 1 {
		if i+1 < length && array[i] < array[i+1] {
			i++
		}
		if array[i] > temp {
			array[pos] = array[i]
			pos = i
		} else {
			break
		}
		array[i] = temp
	}
}

// HeapSort sorts array with Heap Sort Method
func HeapSort(array []int) {
	var temp int
	// Build heap
	for i := len(array)/2 - 1; i >= 0; i-- {
		HeapAdjust(array, i, len(array))
	}
	// Sort
	for j := len(array) - 1; j > 0; j-- {
		temp = array[j]
		array[j] = array[0]
		array[0] = temp
		HeapAdjust(array, 0, j)
	}
}

// QuickSort sorts array with Quick Sort Method
func QuickSort(array []int, left int, right int) {
	if left < right {
		var pivot = array[left]
		var low = left
		var high = right

		for low < high {
			for low < high && array[high] >= pivot {
				high--
			}
			array[low] = array[high]
			for low < high && array[low] <= pivot {
				low++
			}
			array[high] = array[low]
		}
		array[low] = pivot
		QuickSort(array, left, low-1)
		QuickSort(array, low+1, right)
	}
}
