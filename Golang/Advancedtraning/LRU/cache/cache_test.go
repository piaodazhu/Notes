package cache

import "testing"

func TestLRUCache(t *testing.T) {
	c := New[float32](3)
	// c.Print()
	val, ok := c.Get(1)
	if ok || val != nil {
		t.Error("get")
	}

	c.Set(1, "one")
	c.Set(2, "two")
	c.Set(3, "three")

	_, ok = c.Set(3, "Three")
	if !ok {
		t.Errorf("set")
	}

	_, ok = c.Set(4, "four")
	if ok {
		t.Error("set")
	}

	val, ok = c.Get(2)
	if !ok || val != "two" {
		t.Error("get")
	}

	val, ok = c.Del(3)
	if !ok || val != "Three" {
		t.Error("del")
	}

	c.Set(5, "five")
	c.Set(6, "six")

	c.Get(2)
}

func BenchmarkLRUCache(b *testing.B) {
	c := New[int](3)
	for n := 0; n < b.N; n++ {
		c.Set(n, "xxx")
	}
}