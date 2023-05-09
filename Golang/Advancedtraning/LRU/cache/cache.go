package cache

import "fmt"

type KEY interface {
	~int64 | ~int | ~int32 | ~uint | ~uint32 | ~uint64 | string | float32
}

type LRUCache [K KEY] struct {
	nodemap map[K]*cacheNode[K]
	head, tail cacheNode[K]
	cap, size int 
}
type cacheNode [K KEY] struct {
	key	K
	value	any
	prev, next *cacheNode[K]
}

func New[K KEY](capacity int) *LRUCache[K] {
	cache := new(LRUCache[K])
	cache.nodemap = map[K]*cacheNode[K]{}
	cache.cap = capacity
	cache.size = 0
	cache.head.next = &cache.tail
	cache.head.prev = nil
	cache.tail.next = nil
	cache.tail.prev = &cache.head
	// fmt.Printf("%v, %p\n",cache.head, &cache.head)
	// fmt.Printf("%v, %p\n",cache.tail, &cache.tail)
	// fmt.Printf("cache: %p\n", cache)
	return cache
}

func (cache *LRUCache[K]) Init() {
	cache.size = 0
	cache.head.next = &cache.tail
	cache.head.prev = nil
	cache.tail.next = nil
	cache.tail.prev = &cache.head
	cache.nodemap = map[K]*cacheNode[K]{}
}

func (cache *LRUCache[K]) Set(key K, value any) (any, bool) {
	node, ok := cache.nodemap[key]
	if ok {
		oldv := node.value 
		node.value = value
		cache.detachNode(node)
		cache.insertHead(node)
		return oldv, true
	}
	if cache.cap == cache.size {
		replaced := cache.removeTail()
		delete(cache.nodemap, replaced.key)
	}
	newnode := &cacheNode[K]{key: key, value: value}
	cache.nodemap[key] = newnode
	cache.insertHead(newnode)
	return nil, false 
}

func (cache *LRUCache[K]) Get(key K) (any, bool) {
	node, ok := cache.nodemap[key]
	if ok {
		cache.detachNode(node)
		cache.insertHead(node)
		return node.value, true
	}
	return nil, false 
}

func (cache *LRUCache[K]) Del(key K) (any, bool) {
	node, ok := cache.nodemap[key]
	if ok {
		cache.detachNode(node)
		return node.value, true
	}
	return nil, false
}

func (cache *LRUCache[K]) Print() {
	// fmt.Printf("%v, %p\n",cache.head, &cache.head)
	// fmt.Printf("%v, %p\n",cache.tail, &cache.tail)
	
	fmt.Println("------ Print ------")
	p := cache.head.next
	for p != &cache.tail {
		fmt.Printf("key: %v, value: %v\n", p.key, p.value)
		p = p.next
	}
	fmt.Println("------ Finish ------")
}

func (cache *LRUCache[K]) insertHead(node *cacheNode[K]) {
	node.next = cache.head.next
	node.next.prev = node
	cache.head.next = node
	node.prev = &cache.head
	cache.size++
}

func (cache *LRUCache[K]) detachNode(node *cacheNode[K]) {
	node.next.prev = node.prev
	node.prev.next = node.next
	cache.size--
}

func (cache *LRUCache[K]) removeTail() *cacheNode[K] {
	t := cache.tail.prev
	cache.tail.prev = t.prev
	t.prev.next = &cache.tail
	cache.size--
	return t 
}
