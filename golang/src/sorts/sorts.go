// Copyright 2018 Cloud Lam. All rights reserved
// Package sorts implements some sort methods

package sorts

// BubbleSort sorts array with Bubble Sort method
func BubbleSort(array []int) []int {
	for i := 1; i < len(array); i++ {
		for j := 0; j < len(array)-i; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array
}

// InsertSort sorts array with Insert Sort method
func InsertSort(array []int) []int {
	for i := 0; i < len(array); i++ {
		for j := i; j > 0 && array[j-1] > array[j]; j-- {
			array[j-1], array[j] = array[j], array[j-1]
		}
	}
	return array
}

// ShellSort sorts array with Shell's Sort Method
func ShellSort(array []int) []int {
	for gap := len(array) / 2; gap > 0; gap /= 2 {
		for i := gap; i < len(array); i++ {
			if array[i] < array[i-gap] {
				for j := i - gap; j >= 0 && array[j] > array[j+gap]; j -= gap {
					array[j], array[j+gap] = array[j+gap], array[j]
				}
			}
		}
	}
	return array
}

// SelectSort sorts array with Select Sort Method
func SelectSort(array []int) []int {
	var min, max int
	for i := 0; i < len(array)/2; i++ {
		min = i
		max = len(array) - i - 1
		for j := i; j < len(array)-i; j++ {
			if array[j] > array[max] {
				max = j
				continue
			}
			if array[j] < array[min] {
				min = j
			}
		}
		// Minimum
		array[i], array[min] = array[min], array[i]
		// Maximum
		if max == i {
			max = min
		}
		array[len(array)-i-1], array[max] = array[max], array[len(array)-i-1]
	}
	return array
}

// heapAdjust adjust array
// The array must conform to a heap except array[pos],
// adjust array[pos] to make sure array[i] >= array[2i+1] && array[i] >= array[2i+2].
func heapAdjust(array []int, pos int, length int) {
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
func HeapSort(array []int) []int {
	// Build heap
	for i := len(array)/2 - 1; i >= 0; i-- {
		heapAdjust(array, i, len(array))
	}
	// Sort
	for j := len(array) - 1; j > 0; j-- {
		array[j], array[0] = array[0], array[j]
		heapAdjust(array, 0, j)
	}
	return array
}

// QuickSort sorts array with Quick Sort Method
func QuickSort(array []int, left int, right int) []int {
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
	return array
}

// merge array
func merge(array, aux []int, low, mid, high int) {
	for i := low; i <= high; i++ {
		aux[i] = array[i]
	}
	j := mid + 1
	for k := low; k <= high; k++ {
		if low > mid {
			array[k] = aux[j]
			j++
		} else if j > high {
			array[k] = aux[low]
			low++
		} else if aux[low] > aux[j] {
			array[k] = aux[j]
			j++
		} else {
			array[k] = aux[low]
			low++
		}
	}
}

// mergeSortUTD merge array from up to down
func mergeSortUTD(array, aux []int, low, high int) {
	if low >= high {
		return
	}
	mid := (low + high) / 2
	mergeSortUTD(array, aux, low, mid)
	mergeSortUTD(array, aux, mid+1, high)
	merge(array, aux, low, mid, high)
}

// MergeSort sorts array with Merge Sort Method
func MergeSort(array []int) []int {
	aux := make([]int, len(array))
	mergeSortUTD(array, aux, 0, len(array)-1)
	return array
}
