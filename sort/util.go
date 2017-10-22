package sort

func checkArray(array []int) {
	if array == nil {
		panic("Nil array.")
	} else {
		panic("Empty array.")
	}
}

func swap(array []int, index1 int, index2 int) {
	tmp := array[index1]
	array[index1] = array[index2]
	array[index2] = tmp
}
