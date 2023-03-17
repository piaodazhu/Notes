package main

import (
	"myheap/heapimpl"
	"fmt"
)

type Heap struct {
	Items []int
}

func (h Heap) Len() int           { return len(h.Items) }
func (h Heap) Swap(i, j int)      { h.Items[i], h.Items[j] = h.Items[j], h.Items[i] }
func (h Heap) Less(i, j int) bool { return h.Items[i] < h.Items[j] }
func (h *Heap) Push(x any) {
	(*h).Items = append((*h).Items, x.(int))
}
func (h *Heap) Pop() any {
	x := (*h).Items[h.Len()-1]
	(*h).Items = (*h).Items[:h.Len()-1]
	return x
}

func main() {
	array := []int{50, 45, 40, 20, 25, 35, 30, 10, 15}
	h := &Heap{array}
	heap.Init(h)
	heap.Push(h, 99)
	heap.Push(h, 34)
	heap.Push(h, 56)
	heap.Push(h, 8)
	heap.Push(h, 0)
	for h.Len() > 4 {
		x := heap.Pop(h)
		fmt.Println(x)
	}
	fmt.Println("----")
	heap.Push(h, 34)
	heap.Push(h, 56)
	heap.Push(h, 8)
	for h.Len() > 0 {
		x := heap.Pop(h)
		fmt.Println(x)
	}
}
