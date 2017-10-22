package sort

func InsertSort(array []int) {
	checkArray(array)
	// 数组将会被分为两部分一部分是以排序的，
	// 从未排序的部分抽出数据插入已排序的部分
	for i := 1; i < len(array); i++ {
		for j := i; j > 0; j-- {
			if array[j-1] > array[j] {
				swap(array, j-1, j)
			}
		}
	}
}
