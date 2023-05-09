package utils

import (
	"fmt"
	"math/rand"
)

var count int = 0

func init() {
	count = rand.Intn(100)
	fmt.Println("Count is set to ", count)
}

func Count() {
	count++
	fmt.Println("Count", count)
}
