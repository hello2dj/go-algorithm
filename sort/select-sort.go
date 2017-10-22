package sort

func SelectSort(array []int) {
	checkArray(array)

	for i := 0; i < len(array); i++ {
		minIndex := i
		for j := i + 1; j < len(array); j++ {
			if array[j] < array[minIndex] {
				minIndex = j
			}
			if minIndex != i {
				swap(array, i, minIndex)
			}
		}
	}
}
