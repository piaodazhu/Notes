package lfu

import (
	"math/rand"
	"testing"
)

func TestNew(t *testing.T) {
	cache := New[string](3)
	if cache.cap != 3 {
		t.Errorf("cap")
	}
	if cache.mp == nil || len(cache.mp) != 0 {
		t.Errorf("mp")
	}
	if cache.nheap.Len() != 0 {
		t.Error("nheap")
	}
	if cache.size != 0 {
		t.Errorf("size")
	}
	if cache.opcnt != 0 {
		t.Errorf("opcont")
	}
}

func TestSetGetDel(t *testing.T) {
	cache := New[string](3)
	cache.Set("hello", 123)
	x := cache.Get("hello")
	if x.(int) != 123 {
		t.Errorf("Set Get Error")
	}
	if cache.mp["hello"].freq != 1 {
		t.Errorf("Set freq Error, freq")
	}
	if cache.mp["hello"].lastop != 2 {
		t.Errorf("Set lastop Error")
	}
	if cache.mp["hello"].key != "hello" {
		t.Errorf("Set key Error")
	}

	cache.Set("hello", 321)
	x = cache.Get("hello")
	if x.(int) != 321 {
		t.Errorf("Set Get Error2")
	}
	if cache.mp["hello"].freq != 3 {
		t.Errorf("Set freq Error2")
	}
	if cache.mp["hello"].lastop != 4 {
		t.Errorf("Set lastop Error2")
	}

	cache.Del("hello")
	x = cache.Get("hello")
	if x != nil {
		t.Errorf("Del Error")
	}
}

func TestReplace(t *testing.T) {
	cache := New[int](3)
	cache.Set(1, "ONE")
	cache.Get(1)
	cache.Set(1, "one")
	if cache.mp[1].freq != 2 {
		t.Errorf("Wrong freq1")
	}

	cache.Set(2, "two")
	if cache.mp[2].freq != 0 {
		t.Errorf("Wrong freq2")
	}

	cache.ShowHeap()

	cache.Set(3, "Three")
	cache.Get(3)
	cache.Get(3)
	cache.Get(3)
	if cache.mp[3].freq != 3 {
		t.Errorf("Wrong freq3")
	}

	if cache.Get(2).(string) != "two" {
		t.Errorf("Wrong Get")
	}

	cache.ShowHeap()

	cache.Set(4, "Four")
	if _, ok := cache.mp[2]; ok {
		t.Fatalf("Wring Replacement with map, size=%d, cap=%d", cache.size, cache.cap)
	}
	x := cache.Get(2)
	if x != nil {
		t.Errorf("Wrong Replacement with map 2")
	}
	if len(cache.mp) != 3 || cache.nheap.Len() != 3 || cache.size != 3 {
		t.Errorf("Wrong Replacement with size")
	}

	cache.ShowHeap()

	cache.Set(5, "Five")
	x = cache.Get(4)
	if x != nil {
		t.Errorf("Wrong Replacement with map 4")
	}
	if len(cache.mp) != 3 || cache.nheap.Len() != 3 || cache.size != 3 {
		t.Errorf("Wrong Replacement with size 4")
	}

	cache.Get(5)
	if cache.mp[5].freq != 1 {
		t.Errorf("Wrong freq5")
	}

	cache.ShowHeap()

	cache.Set(6, "Six")

	cache.ShowHeap()

	x = cache.Get(5)
	if x != nil {
		t.Errorf("Wrong Replacement with map 5")
	}
	if len(cache.mp) != 3 || cache.nheap.Len() != 3 || cache.size != 3 {
		t.Errorf("Wrong Replacement with size 5")
	}

	cache.Del(1)
	cache.Set(7, "Seven")
	y := cache.Get(1)
	z := cache.Get(7)
	if y != nil || z.(string) != "Seven" {
		t.Errorf("Wrong Del and Set")
	}
}

func BenchmarkSetNew(b *testing.B) {
	cache := New[int](1000)
	for i := 0; i < 100000; i++ {
		cache.Set((i*17)%1000, i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cache.Set((i*107)%b.N+1000, i)
	}
}

func BenchmarkSetOld(b *testing.B) {
	cache := New[int](1000)
	for i := 0; i < 100000; i++ {
		cache.Set((i*17)%1000, i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cache.Set((i*17)%1000, i)
	}
}

func BenchmarkGetNone(b *testing.B) {
	cache := New[int](1000)
	for i := 0; i < 100000; i++ {
		cache.Set((i*17)%1000, i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = cache.Get((i*107)%b.N + 1000)
	}
}

func BenchmarkGet(b *testing.B) {
	cache := New[int](1000)
	for i := 0; i < 100000; i++ {
		cache.Set((i*17)%1000, i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = cache.Get((i * 17) % 1000)
	}
}

func BenchmarkDel(b *testing.B) {
	cache := New[int](b.N * 10)
	for i := 0; i < b.N*10; i++ {
		cache.Set((i*107)%(b.N*10), i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cache.Del((i * 997) % b.N)
	}
}

func BenchmarkMapInsert(b *testing.B) {
	mp := testmp[int]{}
	for i := 0; i < b.N*10; i++ {
		mp.MapInsert(i, i+1)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mp.MapInsert(i+b.N*10, i)
	}
}

func BenchmarkMapDelete(b *testing.B) {
	mp := testmp[int]{}
	for i := 0; i < b.N*10; i++ {
		mp.MapInsert(i, i+1)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mp.MapDelete((i * 997) % (b.N * 10))
	}
}

func BenchmarkHeapInsert(b *testing.B) {
	heap := &nodeheap[int]{}
	nlist := make([]node[int], b.N*10)
	for i := 0; i < b.N*10; i++ {
		nlist[i].freq = rand.Int() % 20
		nlist[i].lastop = i
		nlist[i].index = i
		heap.HeapInsert(&nlist[i])
	}
	newnlist := make([]node[int], b.N)
	for i := 0; i < b.N; i++ {
		newnlist[i].freq = rand.Int() % 20
		newnlist[i].lastop = i
		newnlist[i].index = b.N*10 + i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		heap.HeapInsert(&newnlist[i])
	}
}

func BenchmarkHeapRemove(b *testing.B) {
	heap := &nodeheap[int]{}
	nlist := make([]node[int], b.N*10)
	for i := 0; i < b.N*10; i++ {
		nlist[i].freq = rand.Int() % 20
		nlist[i].lastop = i
		nlist[i].index = i
		heap.HeapInsert(&nlist[i])
	}
	idxspace := make([]int, b.N*10)
	for i := 0; i < b.N*10; i++ {
		idxspace[i] = i
	}
	remlist := make([]int, b.N)
	for i := 0; i < b.N; i++ {
		x := rand.Int()%(b.N*10-i*2) + i
		remlist[i] = idxspace[x]
		idxspace[x] = idxspace[i]
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		heap.HeapRemove(remlist[i])
	}
}

func BenchmarkHeapFix(b *testing.B) {
	heap := &nodeheap[int]{}
	nlist := make([]node[int], b.N*10)
	for i := 0; i < b.N*10; i++ {
		nlist[i].freq = rand.Int() % 20
		nlist[i].lastop = i
		nlist[i].index = i
		heap.HeapInsert(&nlist[i])
	}
	idxspace := make([]int, b.N*10)
	for i := 0; i < b.N*10; i++ {
		idxspace[i] = i
	}
	fixlist := make([]int, b.N)
	for i := 0; i < b.N; i++ {
		x := rand.Int()%(b.N*10-i) + i
		fixlist[i] = idxspace[x]
		idxspace[x] = idxspace[i]
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		nlist[fixlist[i]].freq += (rand.Int()%20 - 10)
		heap.HeapFix(nlist[fixlist[i]].index)
	}
}
