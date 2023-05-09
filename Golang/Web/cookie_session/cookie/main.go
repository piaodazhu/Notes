package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beeku"
)

// avoid duplicate post
var tokenSet map[string]bool

func init() {
	tokenSet = make(map[string]bool)
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析url传递的参数，对于POST则解析响应包的主体（request body）
	//注意:如果没有调用ParseForm方法，下面无法获取表单的数据
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
}

func validSelectValue(s string) bool {
	slice := []string{"apple", "pear", "banane"}
	for _, item := range slice {
		if item == s {
			return true
		}
	}

	return false
}

func validCheckBoxValue(s []string) bool {
	slice := []string{"football", "basketball", "tennis"}
	s1 := make([]interface{}, len(slice))
	for idx, item := range slice {
		s1[idx] = item
	}

	s2 := make([]interface{}, len(s))
	for idx, item := range s {
		s2[idx] = item
	}
	a := beeku.Slice_diff(s2, s1)
	return a == nil
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		// generate timestamp token, to prevent duplicated posting
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		for _, cookie := range r.Cookies() {
			fmt.Println("cookie: ", cookie.Name, cookie.Value)
		}
		if len(r.Cookies()) > 0 {
			interest, _ := r.Cookie("interest")
			fmt.Fprintln(w, "you must like: ", strings.Split(interest.Value, "&"))
			return
		}

		t, _ := template.ParseFiles("login.gtpl")
		log.Println(t.Execute(w, token))
		tokenSet[token] = false
	} else {
		//请求的是登录数据，那么执行登录的逻辑判断
		r.ParseForm()

		// check the token
		token := r.Form.Get("token")
		if token != "" {
			//验证token的合法性
			fmt.Println(token)
			if res, ok := tokenSet[token]; ok {
				if res {
					fmt.Fprintln(w, "duplicate form post!")
					return
				} else {
					tokenSet[token] = true
				}
			} else {
				fmt.Fprintln(w, "invalid token!")
				return
			}
		} else {
			//不存在token报错
			fmt.Fprintln(w, "no token.")
			return
		}

		// XSS vunlerable
		// fmt.Println("username:", r.Form.Get("username"))
		// fmt.Fprintln(w, r.Form.Get("username"))

		// should be like this
		name := template.HTMLEscapeString(r.Form.Get("username"))
		pass := template.HTMLEscapeString(r.Form.Get("password"))
		fmt.Println("username:", name)

		fmt.Println("age:", template.HTMLEscapeString(r.Form.Get("age")))

		fmt.Println("password:", pass)
		if !validSelectValue(r.Form.Get("fruit")) {
			fmt.Println("invalid fruit")
			return
		} else {
			fmt.Println(r.Form["fruit"])
		}

		if !validCheckBoxValue(r.Form["interest"]) {
			fmt.Println("invalid interest")
			return
		} else {
			fmt.Println(r.Form["interest"])
		}
		
		expiration := time.Now().Add(time.Second * 60)
		http.SetCookie(w, &http.Cookie{Name: "fruit", Value: r.Form.Get("fruit"), Expires: expiration})
		http.SetCookie(w, &http.Cookie{Name: "interest", Value: strings.Join(r.Form["interest"], "&"), Expires: expiration})
		fmt.Println("cookie is set")
	}
}

func main() {
	http.HandleFunc("/", sayhelloName) //设置访问的路由
	http.HandleFunc("/login", login)   //设置访问的路由

	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
