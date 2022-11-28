package main

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // 导入数据库驱动
)

// Model Struct
type Userinfo struct {
	Uid        int `orm:"pk;auto"`
	Username   string
	Departname string
}

type User struct {
	Uid     int `orm:"pk"`
	Name    string
	Profile *Profile `orm:"rel(one)"`
	Post    []*Post  `orm:"reverse(many)"`
}

type Profile struct {
	Id   int
	Age  int16
	User *User `orm:"reverse(one)"`
}

type Post struct {
	Id    int
	Title string
	User  *User  `orm:"rel(fk)"`
	Tags  []*Tag `orm:"rel(m2m)"`
}

type Tag struct {
	Id    int
	Name  string
	Posts []*Post `orm:"reverse(many)"`
}

func init() {
	// 设置默认数据库
	orm.RegisterDataBase("default", "mysql", "parallels:mysqlmima@/testDB?charset=utf8", 30)

	// 注册定义的 model
	orm.RegisterModel(new(Userinfo), new(User), new(Profile), new(Post), new(Tag))

	// 创建 table
	orm.RunSyncdb("default", false, true)
}

func main() {
	o := orm.NewOrm()
	var user Userinfo
	user.Username = "Piaodazhu"
	user.Uid = 2
	// Insert
	id, err := o.Insert(&user)
	if err != nil {
		fmt.Println(err)
		panic("insert")
	}
	// delete
	defer o.Delete(&user)

	fmt.Println(id, user.Uid)
	users := []Userinfo{
		{Username: "Lee"},
		{Username: "Wong"},
		{Username: "Chan"},
	}
	// multi-insert
	successNums, err := o.InsertMulti(3, users)
	if err != nil {
		fmt.Println(err)
		panic("insertmulti")
	}

	// delete using raw SQL
	for _, v := range users {
		u := v
		defer o.Raw("delete from userinfo where username = ?", u.Username).Exec()
	}

	fmt.Println(successNums, users[2].Uid)

	// query by giving primary key
	firstuser := Userinfo{Uid: int(id)}
	if o.Read(&firstuser) != nil {
		fmt.Println(err)
		panic("read")
	}

	// update by giving primary key
	firstuser.Username = "P Dazhu"
	num, err := o.Update(&firstuser)
	if err != nil {
		fmt.Println(err)
		panic("update")
	}
	fmt.Println(num, firstuser.Uid)

	searchuser := Userinfo{Uid: 2}
	// searchuser = User{Name: "Chou"}
	err = o.Read(&searchuser)
	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	}
	fmt.Println(searchuser.Uid, searchuser.Username)

	var quser = []*Userinfo{}
	qs := o.QueryTable("userinfo")
	n, err := qs.Filter("uid__between", 2, 4).Limit(3, 0).All(&quser)
	fmt.Println("ret=", n, err)
	if err == nil && n > 0 {
		for i := 0; i < len(quser); i++ {
			fmt.Println(quser[i].Uid, "+", quser[i].Username)
		}
	} else {
		panic("filter")
	}

	qs = o.QueryTable("userinfo")
	n, err = qs.OrderBy("-uid").All(&quser)
	if err == nil && n > 0 {
		for i := 0; i < len(quser); i++ {
			fmt.Println(quser[i].Uid, "+", quser[i].Username)
		}
	} else {
		panic("filter")
	}
}
