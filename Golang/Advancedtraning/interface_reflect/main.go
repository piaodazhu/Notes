package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	age  int
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

func Info(ns NameSayer) {
	switch ns.(type) {
	case Student:
		fmt.Println("a stu")
	case Teacher:
		fmt.Println("a tea")
	}
	t := reflect.TypeOf(ns)
	fmt.Println("It has these methods:")
	for i := 0; i < t.NumMethod(); i++ {
		fmt.Println(t.Method(i).Name)
		if t.Method(i).Name == "SayName" {
			reflect.ValueOf(ns).Method(i).Call(nil)
		}
	}

	if t.Kind() == reflect.Ptr {
		ptr := reflect.ValueOf(ns)
		v := ptr.Elem()
		v.FieldByName("Name").SetString("DEFAULT")
		fmt.Println("name changed")
	}

	fmt.Println()
}

func main() {
	stu := Student{"xiaoming", 12}
	tea := Teacher{"wang"}
	var ns NameSayer
	fmt.Printf("%T:%v, %T:%v, %T:%v\n", stu, stu, tea, tea, ns, ns)
	ns = stu
	fmt.Printf("%T:%v, %T:%v, %T:%v\n", stu, stu, NameSayer(stu), NameSayer(stu), ns, ns)

	t1 := reflect.TypeOf(stu)
	v1 := reflect.ValueOf(stu)

	fmt.Println(t1.NumField(), t1.NumMethod(), v1.Field(0).String(), v1.Method(1), v1.Method(1).Kind())
	v1.Method(1).Call(nil)

	Info(&stu)
	Info(tea)
	Info(ns)
	Info(NameSayer(tea))
}
