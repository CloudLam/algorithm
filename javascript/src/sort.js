/*!
 * sort.js v1.0.0
 * 2017-2018 Cloud Lam
 * Released under the MIT License.
 */

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
