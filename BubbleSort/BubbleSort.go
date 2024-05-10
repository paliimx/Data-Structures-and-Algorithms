package BubbleSort

import "cmp"

func BubbleSort[T cmp.Ordered](array []T) {
	swapCount := 1
	for swapCount > 0 {
		swapCount = 0
		for itemIndex := 0; itemIndex < len(array)-1; itemIndex++ {
			if cmp.Compare(array[itemIndex], array[itemIndex+1]) == 1 {
				array[itemIndex], array[itemIndex+1] = array[itemIndex+1], array[itemIndex]
				swapCount += 1
			}
		}
	}
}
