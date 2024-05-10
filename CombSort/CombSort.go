package CombSort

import "cmp"

func CombSort[T cmp.Ordered](array []T) {
	gapValue := len(array)
	swapCount := 1
	for gapValue >= 1 && swapCount != 0 {
		if gapValue != 1 {
			gapValue = int(float64(gapValue) / float64(1.3))
		}
		swapCount = 0
		firstItem := 0
		secondItem := gapValue
		for secondItem != len(array) {
			if cmp.Compare(array[firstItem], array[secondItem]) == 1 {
				array[firstItem], array[secondItem] = array[secondItem], array[firstItem]
				swapCount += 1
			}
			firstItem += 1
			secondItem += 1
		}
	}
}
