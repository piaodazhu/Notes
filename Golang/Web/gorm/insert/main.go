package main

import (
	"database/sql"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	ID   int64
	Name sql.NullString `gorm:"default:'Unkown'"`
	Age  int64 `gorm:"default:0"`
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

	db.AutoMigrate(&User{})

	// u := User{Name: "qimi", Age: 18}
	// u := User{Name: sql.NullString{String:"", Valid: true}, Age: 19}
	u := User{Name: sql.NullString{String:"qimi", Valid: true}}
	fmt.Println(db.NewRecord(&u)) // yes, it is a new record
	db.Debug().Create(&u)
	fmt.Println(db.NewRecord(&u)) // no, this record is already in the db

}
