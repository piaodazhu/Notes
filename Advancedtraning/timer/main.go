package main

import (
	"fmt"
	"reflect"
	"time"
)
type Student struct {
	Name string
	age int
}
func (this Student) SayName() {
	fmt.Println(this.Name)
}
func (this Student) Age() int {
	return this.age
}
type Teacher struct {
	Name string
}
func (this Teacher) SayName() {
	fmt.Println("Teacher ", this.Name)
}
type NameSayer interface {
	SayName()
}
func main() {
	var x int = 12
	getType := reflect.TypeOf(x)
	fmt.Println(getType.Name(), getType.Kind())
	t := time.NewTimer(time.Second)
	v, ok := <-t.C
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("err")
	}

	v, ok = <-time.After(time.Second)
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("err")
	}

	ch1 := time.After(time.Second * 1)
	ch2 := time.After(time.Second * 2)
	ch3 := time.After(time.Second * 4)
	timeout := time.After(time.Second * 3)
forloop:
	for {
		select {
		case v, ok := <-ch1:
			{
				if ok {
					fmt.Println("ch1:", v)
				} else {
					fmt.Println("err")
				}
			}
		case v, ok := <-ch2:
			{
				if ok {
					fmt.Println("ch2:", v)
				} else {
					fmt.Println("err")
				}
			}
		case v, ok := <-ch3:
			{
				if ok {
					fmt.Println("ch3:", v)
				} else {
					fmt.Println("err")
				}
			}
		case <-timeout:
			{
				fmt.Println("3 second timeout!")
				break forloop
			}
		}
	}
	fmt.Println("done")
}
