package TernarySearch

import "cmp"

func TernarySearch[T cmp.Ordered](array []T, number T) int {
	minIndex := 0
	maxIndex := len(array) - 1
	for minIndex <= maxIndex {
		midIndex1 := minIndex + int((maxIndex-minIndex)/3)
		midIndex2 := maxIndex - int((maxIndex-minIndex)/3)
		midItem1 := array[midIndex1]
		midItem2 := array[midIndex2]
		if cmp.Compare(midItem1, number) == 0 {
			return midIndex1
		} else if cmp.Compare(midItem2, number) == 0 {
			return midIndex2
		}
		if cmp.Compare(midItem1, number) == -1 {
			minIndex = midIndex1 + 1
		} else if cmp.Compare(midItem2, number) == 1 {
			maxIndex = midIndex2 - 1
		} else {
			minIndex = midIndex1 + 1
			maxIndex = midIndex2 - 1
		}
	}
	return -1
}
