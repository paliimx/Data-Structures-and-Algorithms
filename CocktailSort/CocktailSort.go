package CocktailSort

import "cmp"

func CocktailSort[T cmp.Ordered](array []T) {
	swapCount := 1
	for swapCount > 0 {
		swapCount = 0
		for itemIndex := 0; itemIndex < len(array)-1; itemIndex++ {
			if cmp.Compare(array[itemIndex], array[itemIndex+1]) == 1 {
				array[itemIndex], array[itemIndex+1] = array[itemIndex+1], array[itemIndex]
				swapCount += 1
			}
		}
		for itemIndex := len(array) - 1; itemIndex > 0; itemIndex-- {
			if cmp.Compare(array[itemIndex], array[itemIndex-1]) == -1 {
				array[itemIndex], array[itemIndex-1] = array[itemIndex-1], array[itemIndex]
				swapCount += 1
			}
		}
	}
}
