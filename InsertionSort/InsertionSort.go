package InsertionSort

import "cmp"

func InsertionSort[T cmp.Ordered](array []T) {
	for itemIndex, itemValue := range array {
		for itemIndex != 0 && cmp.Compare(array[itemIndex-1], itemValue) == 1 {
			array[itemIndex] = array[itemIndex-1]
			itemIndex -= 1
		}
		array[itemIndex] = itemValue
	}
}
