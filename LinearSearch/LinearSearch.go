package LinearSearch

import "cmp"

func LinearSearch[T cmp.Ordered](array []T, number T) int {
	for index, value := range array {
		if cmp.Compare(value, number) == 0 {
			return index
		}
	}
	return -1
}
