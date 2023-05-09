package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(time.Second)
	t1 := time.Now()
	fmt.Printf("%T\n%v\n%#v\n", t1, t1, t1)

	t2 := time.Date(2022, 10, 29, 12, 12, 42, 8000, time.Local)
	fmt.Printf("%T\n%v\n%#v\n", t2, t2, t2)

	fmt.Println(t1.Sub(t2))

	fmt.Println(t1.Format("2006nian, 1yue, 2ri"))
	s := "01月01日 2023"
	t3, err := time.Parse("01月02日 2006", s)
	if err != nil {
		panic("time.Parse")
	}
	fmt.Println(t3)

	year, month, day := t1.Date()
	fmt.Println(year, month, day)

	hour, minite, second := t1.Clock()
	fmt.Println(hour, minite, second)

	rand.Seed(time.Now().Unix())
	time.Sleep(time.Second * time.Duration(rand.Intn(10)))

	fmt.Println(t2.Unix())
	fmt.Println(t2.UnixNano())

}