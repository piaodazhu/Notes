package main

import (
	"database/sql"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	Name         string        `gorm:"default:'Unknown'"`
	Age          sql.NullInt64 `gorm:"column:user_age"` // 零值类型
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"`
	Role         string  `gorm:"size:255"`
	MemberNumber *string `gorm:"unique;not null"`
	Num          int     `gorm:"AUTO_INCREMENT"`
	Address      string  `gorm:"index:addr"` // create index of addr
	IgnoreMe     int     `gorm:"-"`          // ignored
}

func (User) TableName() string {
	return "profiles"
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

	// create temporary table
	db.Table("tempUsers").CreateTable(&User{})
}
