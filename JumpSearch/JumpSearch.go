package JumpSearch

import (
	"cmp"
	"math"
)

func JumpSearch[T cmp.Ordered](array []T, number T) int {
	jumpValue := int(math.Floor(math.Sqrt(float64(len(array)))))
	minIndex := 0
	maxIndex := jumpValue
	for cmp.Compare(array[maxIndex], number) == -1 {
		minIndex += jumpValue
		maxIndex = minIndex + jumpValue
		if maxIndex >= len(array) {
			maxIndex = len(array)
			minIndex = maxIndex - jumpValue
			break
		}
	}
	for i := minIndex; i < maxIndex; i++ {
		if cmp.Compare(array[i], number) == 0 {
			return i
		}
	}
	return -1
}
