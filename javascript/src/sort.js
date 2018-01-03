/*!
 * sort.js v1.0.0
 * 2017-2018 Cloud Lam
 * Released under the MIT License.
 */

// Bubble Sort
function bubbleSort (array) {
  for (var i = 1; i < array.length; i++) {
    for (var j = 0; j < array.length - i; j++) {
      if (array[j] > array[j + 1]) {
        var temp = array[j];
        array[j] = array[j + 1];
        array[j + 1] = temp;
      }
    }
  }
}

// Insert Sort
function insertSort (array) {
  for (var i = 1; i < array.length; i++) {
    if (array[i] < array[i - 1]) {
      var temp = array[i];
      for (var j = i - 1; j >= 0 && array[j] > temp; j--) {
        array[j + 1] = array[j];
        array[j] = temp;
      }
    }
  }
}

// Shell Sort
function shellSort (array) {
  for (var gap = array.length / 2; gap > 0; gap /= 2) {
    gap = parseInt(gap);
    for (var i = gap; i < array.length; i++) {
      if (array[i] < array[i - gap]) {
        for (var j = i - gap; j >= 0 && array[j] > array[j + gap]; j -= gap) {
          var temp = array[j];
          array[j] = array[j + gap];
          array[j + gap] = temp;
        }
      }
    }
  }
}

// Select Sort
function selectSort (array) {
  var max, min, temp;
  for (var i = 0; i < array.length / 2; i++) {
    max = i;
    min = i;
    for (var j = i + 1; j < array.length - i; j++) {
      if (array[j] > array[max]) {
        max = j;
        continue;
      }
      if (array[j] < array[min]) {
        min = j;
      }
    }
    // Minimum
    if (min != i) {
      temp = array[i];
      array[i] = array[min];
      array[min] = temp;
    }
    // Maximum
    if (max != i) {
      temp = array[array.length - i - 1];
      array[array.length - i - 1] = array[max];
      array[max] = temp;
    }
  }
}

// Heap Sort
/**
 * Adjust heap
 * The array must conform to a heap except array[pos], 
 * adjust array[pos] to make sure array[i] >= array[2i+1] && array[i] >= array[2i+2].
 * 
 * @param array
 * @param pos
 * @param length
 */
function heapAdjust (array, pos, length) {
  var temp = array[pos];
  for (var i = pos * 2 + 1; i < length; i = i * 2 + 1) {
    if (i + 1 < length && array[i] < array[i + 1]) {
      i++;
    }
    if (array[i] > temp) {
      array[pos] = array[i];
      pos = i;
    } else {
      break;
    }
    array[i] = temp;
  }
}

function heapSort (array) {
  // Build heap
  for (var i = parseInt(array.length / 2 - 1); i >= 0; i--) {
    heapAdjust(array, i, array.length);
  }
  // Sort
  for (var j = array.length - 1; j > 0; j--) {
    var temp = array[j];
    array[j] = array[0];
    array[0] = temp;
    heapAdjust(array, 0, j);
  }
}

// Quick Sort
function quickSort (array, left, right) {
  if (left < right) {
    var pivot = array[left];
    var low = left;
    var high = right;
    while (low < high) {
      while (low < high && array[high] >= pivot) {
        high--;
      }
      array[low] = array[high];
      while (low < high && array[low] <= pivot) {
        low++;
      }
      array[high] = array[low];
    }
    array[low] = pivot;
    quickSort(array, left, low - 1);
    quickSort(array, low + 1, right);
  }
}
