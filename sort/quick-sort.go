package sort

func partitionRecursion(array []int, left int, right int) {
	low := left
	high := right

	// 选择pivot
	pivot := array[(left+right)>>1]
	for low <= high {
		for array[low] < pivot {
			low++
		}
		for array[high] > pivot {
			high--
		}
		if low <= high {
			swap(array, low, high)
			low++
			high--
		}
	}

	if left < high {
		partitionRecursion(array, left, high)
	}

	if low < right {
		partitionRecursion(array, low, right)
	}
}

func QuickSort(array []int) {
	checkArray(array)
	partitionRecursion(array, 0, len(array)-1)
}
