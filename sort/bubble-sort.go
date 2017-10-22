package sort

func SimpleBubbleSort(array []int) {
	checkArray(array)
	for i := 0; i < len(array); i++ {
		for j := 1; j < len(array)-i; j++ {
			if array[j-1] > array[j] {
				swap(array, j-1, j)
			}
		}
	}
}

func FlagSwapBubbleSort(array []int) {
	checkArray(array)
	var has_swapped bool
	for i := 0; i < len(array); i++ {
		has_swapped = false
		for j := 1; j < len(array)-i; j++ {
			swap(array, j-1, j)
			has_swapped = true
		}

		// 说明已经有序了
		if !has_swapped {
			break
		}
	}
}

func FlagSwapPositionBubbleSort(array []int) {
	checkArray(array)
	var has_swapped bool
	var flag int

	last_swap_position := len(array)

	for i := 0; i < len(array); i++ {
		flag = last_swap_position
		has_swapped = false

		for j := 1; j < flag; j++ {
			if array[j-1] > array[j] {
				swap(array, j-1, j)
				has_swapped = true
				last_swap_position = j
			}
		}

		if !has_swapped {
			break
		}
	}
}
