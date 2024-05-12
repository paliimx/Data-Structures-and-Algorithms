package MergeSort

import "cmp"

func mergeParts[T cmp.Ordered](array []T, leftIndex, divideIndex, rightIndex int) {
	var tempArray1, tempArray2 []T
	for i := leftIndex; i <= divideIndex; i++ {
		tempArray1 = append(tempArray1, array[i])
	}
	for i := divideIndex + 1; i <= rightIndex; i++ {
		tempArray2 = append(tempArray2, array[i])
	}
	arrayIndex := leftIndex
	tempArray1Index := 0
	tempArray2Index := 0
	for tempArray1Index != len(tempArray1) && tempArray2Index != len(tempArray2) {
		if tempArray1[tempArray1Index] <= tempArray2[tempArray2Index] {
			array[arrayIndex] = tempArray1[tempArray1Index]
			tempArray1Index += 1
		} else {
			array[arrayIndex] = tempArray2[tempArray2Index]
			tempArray2Index += 1
		}
		arrayIndex += 1
	}
	for tempArray1Index < len(tempArray1) {
		array[arrayIndex] = tempArray1[tempArray1Index]
		tempArray1Index += 1
		arrayIndex += 1

	}
	for tempArray2Index < len(tempArray2) {
		array[arrayIndex] = tempArray2[tempArray2Index]
		tempArray2Index += 1
		arrayIndex += 1
	}
}

func MergeSort[T cmp.Ordered](array []T, leftIndex, rightIndex int) {
	if leftIndex >= rightIndex {
		return
	}
	divideIndex := int((leftIndex + rightIndex) / 2)
	MergeSort(array, leftIndex, divideIndex)
	MergeSort(array, divideIndex+1, rightIndex)
	mergeParts(array, leftIndex, divideIndex, rightIndex)
}
