package heap // 小顶堆

type BinaryHeap struct {
	array []int
}

func (heap *BinaryHeap) Size() int {
	return len(heap.array) - 1
}

func (heap *BinaryHeap) Add(data int) {
	if len(heap.array) == 0 {
		heap.array = append(heap.array, -1) //哨兵
	}
	heap.array = append(heap.array, data)
	for i := heap.Size(); heap.array[i] < heap.array[i/2]; i /= 2 {
		tmp := heap.array[i]
		heap.array[i] = heap.array[i/2]
		heap.array[i/2] = tmp
	}
}

func (heap *BinaryHeap) RemoveMinimum() int {
	if heap.Size() == 0 {
		panic("Empty heap.")
	}
	minimum := heap.array[1]
	last := heap.array[heap.Size()]

	var i int
	var child int
	for i = 1; i*2 < heap.Size(); i = child { // 重新构建堆
		child = i * 2                                                       // 左孩子
		if child < heap.Size() && heap.array[child] > heap.array[child+1] { // 取小的
			child++
		}
		if last > heap.array[child] {
			tmp := heap.array[i]
			heap.array[i] = heap.array[child]
			heap.array[child] = tmp
		} else {
			break
		}
	}
	heap.array[i] = last //存储last
	newArray := heap.array[0:i]
	heap.array = newArray
	return minimum
}
