package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserInfo struct {
	ID     uint
	Name   string
	Gender string
	Hobby  string
}

func ErrProc(err error, tag string) {
	if err == nil {
		return
	}
	panic("[" + tag + "]" + err.Error())
}

func main() {
	db, err := gorm.Open("mysql", "parallels:mysqlmima@(localhost)/testDB?charset=utf8&parseTime=True")
	ErrProc(err, "1")
	defer db.Close()

	db.AutoMigrate(&UserInfo{})

	u1 := UserInfo{1, "Lee", "M", "Music"}
	db.Create(&u1)

	var u UserInfo
	db.First(&u)
	fmt.Println(u)

	db.Model(&u).Update("hobby", "swim")
	db.First(&u)
	fmt.Println(u)

	db.Delete(&u)
}
