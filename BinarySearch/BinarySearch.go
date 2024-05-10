package BinarySearch

import "cmp"

func BinarySearch[T cmp.Ordered](array []T, value T) int {
	minIndex := 0
	maxIndex := len(array) - 1
	for minIndex <= maxIndex {
		midIndex := (maxIndex + minIndex) / 2
		midItem := array[midIndex]
		if value == midItem {
			return midIndex
		}
		if cmp.Less(midItem, value) {
			minIndex = midIndex + 1
		} else {
			maxIndex = midIndex - 1
		}
	}
	return -1
}
