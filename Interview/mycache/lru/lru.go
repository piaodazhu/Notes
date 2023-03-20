package lru

type KeyType interface {
	~int | ~uint | ~string | ~float32 | ~float64
}

type LRU[T KeyType] struct {
	size int
	cap  int
	mp   map[T]*node[T]
	head *node[T]
	tail *node[T]
}

type node[T KeyType] struct {
	key T
	value interface{}
	next *node[T]
	prev *node[T]
}

func New[T KeyType](cap int) *LRU[T] {
	lru := &LRU[T]{
		size: 0,
		cap: cap,
		mp: map[T]*node[T]{},
		head: &node[T]{
			next: nil,
			prev: nil,
		},
		tail: &node[T]{
			next: nil,
			prev: nil,
		},
	}
	lru.head.next = lru.tail
	lru.tail.prev = lru.head
	return lru
}

func (c *LRU[T]) Set(key T, value interface{}) {
	if n, ok := c.mp[key]; ok {
		detach(n)
		insert(c.head.next, n)
		return
	}
	n := &node[T]{
		key: key,
		value: value,
	}
	c.mp[key] = n
	insert(c.head.next, n)
	if c.size < c.cap {
		c.size++
	} else {
		kick := c.tail.prev
		detach(kick)
		delete(c.mp, kick.key)
	}
}

func (c *LRU[T]) Get(key T) (interface{}, bool) {
	if n, ok := c.mp[key]; ok {
		detach(n)
		insert(c.head.next, n)
		return n.value, true
	}	
	return nil, false
}

func (c *LRU[T]) Del(key T) {
	if n, ok := c.mp[key]; ok {
		detach(n)
		delete(c.mp, n.key)
	}
}

func detach[T KeyType](n *node[T]) {
	n.next.prev = n.prev
	n.prev.next = n.next
}

func insert[T KeyType](place *node[T], n *node[T]) {
	n.next = place
	n.prev = place.prev
	place.prev.next = n
	place.prev = n
}