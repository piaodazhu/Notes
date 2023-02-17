// 方法的接收者只能是结构类型 不能是接口
// 所以绑定在Shape类型上的Area方法，Rectangle如果调用，传入的参数在方法内部还是会当成Shape来处理，而不是Rectangle类
// 在Go做面向对象时应谨慎，它是has a 而不是 is a，父类对象没法调用子类对象上的方法
package main

import (
	"fmt"
	"math"
)

type ShapeInterface interface {
	Area() float64
	GetName() string
	PrintArea()
}

// 标准形状，它的面积为0.0
type Shape struct {
	name string
}

func (s *Shape) Area() float64 {
	return 0.0
}

func (s *Shape) GetName() string {
	return s.name
}

func (s *Shape) PrintArea() {
	// wont work
	fmt.Printf("%s : Area %v\r\n", s.name, s.Area())

	// wont work, either
	// t := reflect.TypeOf(s).Elem()
	// fmt.Println(t.Kind(), t.Name(), t.NumMethod())
	// v := reflect.ValueOf(s)
	// r := v.MethodByName("Area").Call(nil)
	// fmt.Printf("%s : Area %v\r\n", s.name, r[0].Float())
}

// // it will work
func (r *Rectangle) PrintArea() {
	// wont work
	fmt.Printf("%s : Area %v\r\n", r.name, r.Area())
}

// 矩形 : 重新定义了Area方法
type Rectangle struct {
	Shape
	w, h float64
}

func (r *Rectangle) Area() float64 {
	return r.w * r.h
}

func (r *Rectangle) GetName() string {
	return "--->" + r.Shape.name
}

// 圆形  : 重新定义 Area 和PrintArea 方法
type Circle struct {
	Shape
	r float64
}

func (c *Circle) Area() float64 {
	return c.r * c.r * math.Pi
}

func (c *Circle) PrintArea() {
	fmt.Printf("%s : Area %v\r\n", c.GetName(), c.Area())
}

func main() {

	s := Shape{name: "Shape"}
	c := Circle{Shape: Shape{name: "Circle"}, r: 10}
	r := Rectangle{Shape: Shape{name: "Rectangle"}, w: 5, h: 4}

	listshape := []ShapeInterface{&s, &c, &r}

	for _, si := range listshape {
		si.PrintArea() //!! 猜猜哪个Area()方法会被调用 !!
		fmt.Println(si.GetName())
	}

}
