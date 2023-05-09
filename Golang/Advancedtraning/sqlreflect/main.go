package main

import (
	"fmt"
	"reflect"
)

type order struct {
	ordId      int
	CustomerId int `tag1:"123" tag2:"str"`
	beizhu     int
}

type employee struct {
	name    string
	id      int
	address string
	salary  int
	country string
}

func createQuery(q interface{}) string {
	t := reflect.TypeOf(q)
	v := reflect.ValueOf(q)
	// fmt.Println(v.Kind(), t.Kind())
	if v.Kind() != reflect.Struct {
		panic("unsupported argument type!")
	}
	tableName := t.Name() // 通过结构体类型提取出SQL的表名
	sql := fmt.Sprintf("INSERT INTO %s ", tableName)
	columns := "("
	values := "VALUES ("
	// fmt.Println(t.NumField(), v.NumField())
	// fmt.Println(t.Field(1).Index, t.Field(1).Offset, t.Field(1).IsExported())
	// fmt.Println(v.Field(0).CanInt(), v.Field(0).CanInterface())
	// fmt.Println(t.Field(1).Tag.Get("tag1"))
	for i := 0; i < v.NumField(); i++ {
		// 注意reflect.Value 也实现了NumField,Kind这些方法
		// 这里的v.Field(i).Kind()等价于t.Field(i).Type.Kind()
		switch v.Field(i).Kind() {
		case reflect.Int:
			if i == 0 {
				columns += fmt.Sprintf("%s", t.Field(i).Name)
				values += fmt.Sprintf("%d", v.Field(i).Int())
			} else {
				columns += fmt.Sprintf(", %s", t.Field(i).Name)
				values += fmt.Sprintf(", %d", v.Field(i).Int())
			}
		case reflect.String:
			if i == 0 {
				columns += fmt.Sprintf("%s", t.Field(i).Name)
				values += fmt.Sprintf("'%s'", v.Field(i).String())
			} else {
				columns += fmt.Sprintf(", %s", t.Field(i).Name)
				values += fmt.Sprintf(", '%s'", v.Field(i).String())
			}
		}
	}
	columns += "); "
	values += "); "
	sql += columns + values
	fmt.Println(sql)
	return sql
}

func main() {
	o := order{
		ordId:      456,
		CustomerId: 56,
	}
	createQuery(o)

	e := employee{
		name:    "Naveen",
		id:      565,
		address: "Coimbatore",
		salary:  90000,
		country: "India",
	}
	createQuery(e)
}
