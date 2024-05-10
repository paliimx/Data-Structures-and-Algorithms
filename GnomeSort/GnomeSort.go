package GnomeSort

import "cmp"

func GnomeSort[T cmp.Ordered](array []T) {
	itemIndex := 0
	for itemIndex < len(array)-1 {
		if cmp.Compare(array[itemIndex], array[itemIndex+1]) == 1 {
			array[itemIndex], array[itemIndex+1] = array[itemIndex+1], array[itemIndex]
			if itemIndex != 0 {
				itemIndex -= 1
			}
		} else {
			itemIndex += 1
		}
	}
}
