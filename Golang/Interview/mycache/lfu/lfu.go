package lfu

import (
	"container/heap"
)

type KeyType interface {
	~int | ~string | ~uint | ~float32 | ~float64
}

type LFU[T KeyType] struct {
	size  int
	cap   int
	mp    map[T]*node[T]
	nheap nodeheap[T]
	opcnt int
}

func (c *LFU[T]) GetOp() int {
	c.opcnt++
	return c.opcnt
}

type nodeheap[T KeyType] []*node[T]

func (h nodeheap[T]) Len() int { return len(h) }
func (h nodeheap[T]) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].index, h[j].index = h[j].index, h[i].index
}
func (h nodeheap[T]) Less(i, j int) bool {
	if h[i].freq < h[j].freq {
		return true
	} else if h[i].freq > h[j].freq {
		return false
	} else {
		return h[i].lastop < h[j].lastop
	}
}

func (h *nodeheap[T]) Push(x interface{}) {
	(*h) = append((*h), x.(*node[T]))
}
func (h *nodeheap[T]) Pop() interface{} {
	x := (*h)[h.Len()-1]
	(*h) = (*h)[:h.Len()-1]
	return x
}

type node[T KeyType] struct {
	key    T
	value  interface{}
	freq   int
	lastop int
	index  int
}

func New[T KeyType](cap int) *LFU[T] {
	lfu := &LFU[T]{
		size:  0,
		cap:   cap,
		mp:    map[T]*node[T]{},
		nheap: nodeheap[T]{},
		opcnt: 0,
	}
	return lfu
}

func (c *LFU[T]) Set(key T, value interface{}) {
	if n, ok := c.mp[key]; ok {
		n.value = value
		n.freq++
		n.lastop = c.GetOp()
		heap.Fix(&c.nheap, n.index)
	} else {
		newnode := &node[T]{
			key:    key,
			value:  value,
			freq:   0,
			lastop: c.GetOp(),
			index:  c.nheap.Len() - 1,
		}

		if c.size == c.cap {
			kicked := heap.Pop(&c.nheap)
			delete(c.mp, kicked.(*node[T]).key)
		} else {
			c.size++
			newnode.index++
		}
		heap.Push(&c.nheap, newnode)
		c.mp[key] = newnode
	}
}

func (c *LFU[T]) Get(key T) interface{} {
	if n, ok := c.mp[key]; ok {
		n.freq++
		n.lastop = c.GetOp()
		heap.Fix(&c.nheap, n.index)
		return n.value
	} else {
		return nil
	}
}

func (c *LFU[T]) Del(key T) {
	if n, ok := c.mp[key]; ok {
		heap.Remove(&c.nheap, n.index)
		c.size--
		delete(c.mp, n.key)
	}
}

func (c *LFU[T]) ShowHeap() {
	// fmt.Println("------HEAP------")
	// for i := range c.nheap {
	// 	fmt.Printf("%+v\n", c.nheap[i])
	// }
	// fmt.Println("")
}

type testmp[T KeyType] map[T]interface{}

func (tmp testmp[T])MapInsert(key T, value interface{}) {
	tmp[key] = value
}

func (tmp testmp[T])MapDelete(key T) {
	delete(tmp, key)
}

func (hp *nodeheap[T])HeapInsert(n *node[T]) {
	heap.Push(hp, n)
}

func (hp *nodeheap[T])HeapRemove(i int) {
	heap.Remove(hp, i)
}

func (hp *nodeheap[T])HeapFix(i int) {
	heap.Fix(hp, i)
}