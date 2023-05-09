package main

import (
	"test/cache"
)

func main() {
	c := cache.New[float32](3)
	c.Print()
	val, ok := c.Get(1)
	if ok || val != nil {
		panic("get")
	}

	c.Set(1, "one")
	c.Set(2, "two")
	c.Set(3, "three")
	c.Print()

	_, ok = c.Set(3, "Three")
	if !ok {
		panic("set")
	}
	c.Print()

	_, ok = c.Set(4, "four")
	if ok {
		panic("set")
	}

	c.Print()

	val, ok = c.Get(2)
	if !ok || val != "two" {
		panic("get")
	}

	c.Print()

	val, ok = c.Del(3)
	if !ok || val != "Three" {
		panic("del")
	}

	c.Print()

	c.Set(5, "five")
	c.Set(6, "six")
	c.Print()
	c.Get(2)
	c.Print()
}
