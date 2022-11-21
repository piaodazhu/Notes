package main

import (
	"fmt"
	"reflect"
)

type Card struct {
	Id    string
	IdNum uint64
}

type Student struct {
	Name    string
	Age     int8
	Height  float64
	StdCard Card
}

func (s Student) Eat(food string) {
	fmt.Printf("%s eat %s\n", s.Name, food)
}

func reflect_value(a interface{}) {
	v := reflect.ValueOf(a)
	fmt.Println(v)
	k := v.Kind()
	fmt.Println(k)
	switch k {
	case reflect.Float64:
		fmt.Println("a是 ", v.Float())
	}
}

//反射获取interface类型信息

func reflect_type(a interface{}) {
	t := reflect.TypeOf(a)
	fmt.Println("类型是：", t)
	// kind()可以获取具体类型
	k := t.Kind()
	fmt.Println(k)
	switch k {
	case reflect.Float64:
		fmt.Printf("a is float64\n")
	case reflect.String:
		fmt.Println("string")
	}
}

// 反射修改值
func reflect_set_value(a interface{}) {
	v := reflect.ValueOf(a)
	fmt.Println(v)
	k := v.Kind()
	fmt.Println(k)
	switch k {
	case reflect.Float64:
		panic("you should input a pointer")
	case reflect.Ptr:
		// Elem()获取地址指向的值
		v.Elem().SetFloat(7.9)
		fmt.Println("case:", v.Elem().Float())
		// 地址
		fmt.Println(v.Pointer())
	}
}

func Poni(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("类型：", t)
	fmt.Println("字符串类型：", t.Name())
	// 获取值
	v := reflect.ValueOf(o)
	fmt.Println(v)
	// 可以获取所有属性
	// 获取结构体字段个数：t.NumField()
	for i := 0; i < t.NumField(); i++ {
		// 取每个字段
		f := t.Field(i)
		fmt.Printf("%s : %v", f.Name, f.Type)
		// 获取字段的值信息
		// Interface()：获取字段对应的值
		val := v.Field(i).Interface()
		fmt.Println("val :", val)
	}
	fmt.Println("=================方法====================")
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Println(m.Name)
		fmt.Println(m.Type)
	}

}

func main() {
	// var x float64 = 3.4
	ming := Student{"xiaoming", 18, 168.33, Card{"1234", 1234}}
	ming.Eat("banana")
	reflect_type(ming.StdCard.Id)
	reflect_value(ming.StdCard.IdNum)
	reflect_set_value(&ming.Height)
	fmt.Printf("%#v\n", ming)
	Poni(ming)
}
