package heap

type Myheaper interface {
	Len() int
	Swap(int, int)
	Less(int, int) bool
	Push(any)
	Pop() any
}

func Init(heap Myheaper) {
	for i := (heap.Len() >> 1) - 1; i >= 0; i-- {
		heapify(heap, i)
	}
}

func Push(heap Myheaper, item any) {
	heap.Push(item)
	for i := (heap.Len() >> 1) - 1; i >= 0; i = (i - 1) >> 1 {
		heapify(heap, i)
	}
}

func Pop(heap Myheaper) any {
	heap.Swap(0, heap.Len() - 1)
	tmp := heap.Pop()
	heapify(heap, 0)
	return tmp
}


func heapify(heap Myheaper, root int) {
	for root < heap.Len() {
		left := ((root + 1) << 1) - 1
		right := (root + 1) << 1
		smallest := root
		if right < heap.Len() && heap.Less(smallest, left) {
			smallest = left 
		}
		if right < heap.Len() && heap.Less(smallest, right) {
			smallest = right
		}
		if smallest != root {
			heap.Swap(root, smallest)
			root = smallest
		} else {
			break
		}
	}
}