package ExponentialSearch

import "cmp"

func ExponentialSearch[T cmp.Ordered](array []T, val T) int {
	boundValue := 1
	for boundValue < len(array) && array[boundValue] < val {
		boundValue *= 2
	}
	if boundValue > len(array) {
		boundValue = len(array) - 1
	}
	return BinarySearch(array, boundValue+1, val)
}

func BinarySearch[T cmp.Ordered](array []T, bound int, number T) int {
	minIndex := 0
	maxIndex := bound - 1
	for minIndex <= maxIndex {
		midIndex := int((maxIndex + minIndex) / 2)
		midItem := array[midIndex]
		if cmp.Compare(number, midItem) == 0 {
			return midIndex
		}
		if cmp.Compare(midItem, number) == -1 {
			minIndex = midIndex + 1
		} else {
			maxIndex = midIndex - 1
		}
	}
	return -1
}
