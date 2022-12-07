package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	ID           int64         `gorm:"auto_increment"`
	Name         string        `gorm:"not null"`
	Age          sql.NullInt64
	Birthday     *time.Time
	Email        string `gorm:"type:varchar(100);unique_index"`
	Role         string `gorm:"size:255"`
	MemberNumber string `gorm:"unique;not null"`
	Address      string `gorm:"index:addr"` // create index of addr
	IgnoreMe     int    `gorm:"-"`          // ignored
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

	now := time.Now()
	n1 := User{Name: "Alice", Age: sql.NullInt64{12, true}, Birthday: &now, Email: "123@aa.bb", MemberNumber: "N1", Address: "Beijing"}
	n2 := User{Name: "Bob", Age: sql.NullInt64{14, true}, Birthday: &now, Email: "223@aa.bb", MemberNumber: "N2", Address: "Shanghai"}
	n3 := User{Name: "Chen", Age: sql.NullInt64{18, true}, Birthday: &now, Email: "323@aa.bb", MemberNumber: "N3", Address: "Guangzhou"}
	n4 := User{Name: "Dannel", Age: sql.NullInt64{22, true}, Birthday: &now, Email: "423@aa.bb", MemberNumber: "N4", Address: "Chengdu"}
	n5 := User{Name: "Emma", Age: sql.NullInt64{11, true}, Birthday: &now, Email: "523@aa.bb", MemberNumber: "N5", Address: "Hangzhou"}

	db.Create(&n1)
	defer func() {
		if !db.NewRecord(&n1) {
			db.Delete(&n1)
		}
	}()

	db.Create(&n2)
	defer func() {
		if !db.NewRecord(&n2) {
			db.Delete(&n2)
		}
	}()

	db.Create(&n3)
	defer func() {
		if !db.NewRecord(&n3) {
			db.Delete(&n3)
		}
	}()

	db.Create(&n4)
	defer func() {
		if !db.NewRecord(&n4) {
			db.Delete(&n4)
		}
	}()

	db.Create(&n5)
	defer func() {
		if !db.NewRecord(&n5) {
			db.Delete(&n5)
		}
	}()

	var user User
	db.First(&user)
	fmt.Println("get first user: ", user)

	user = User{}
	db.First(&user, 3)
	fmt.Println("get user which id=3", user)

	user = User{}
	db.Last(&user)
	fmt.Println("get last user: ", user)
	
	user = User{}
	db.Take(&user)
	fmt.Println("randomly pick a user: ", user)

	var users []User
	db.Find(&users)
	fmt.Println("all users:")
	for _, u := range users {
		fmt.Println(u)
	}

	var elderusers []User
	db.Limit(1).Offset(1).Where("age >= ? AND email like ?", "18", "%@aa.bb%").Find(&elderusers)
	fmt.Println("elder users:")
	for _, u := range elderusers {
		fmt.Println(u)
	}

	type Result struct {
		Name string
		Age int
	}

	var results []Result
	fmt.Println("select:")
	db.Table("users").Select("name, age").Order("name DESC").Scan(&results)
	for _, r := range results {
		fmt.Println(r)
	}
}
