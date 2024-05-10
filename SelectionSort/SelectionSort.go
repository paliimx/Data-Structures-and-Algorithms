package SelectionSort

import "cmp"

func SelectionSort[T cmp.Ordered](array []T) {
	for arrayIndex := range array {
		minValue := array[arrayIndex]
		minIndex := arrayIndex
		for subArrayIndex := arrayIndex + 1; subArrayIndex < len(array); subArrayIndex++ {
			if cmp.Compare(array[subArrayIndex], minValue) == -1 {
				minValue = array[subArrayIndex]
				minIndex = subArrayIndex
			}
		}
		array[minIndex], array[arrayIndex] = array[arrayIndex], array[minIndex]
	}
}
