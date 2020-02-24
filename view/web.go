package view

import (
	"SimpleComment/view"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func showcomment(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	res,err := view.Getcomment()
	if err!=nil{
		panic(err)
	}
	for _,s:=range res {
		fmt.Fprintln(w, "用户名："+s.Username)
		fmt.Fprintln(w, "评论内容："+s.Content)
		fmt.Fprintln(w, strconv.Itoa(int(s.Id)))
		fmt.Fprintln(w, "发表时间："+s.Time)
	}
}

func comment_web(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		t, _ := template.ParseFiles("comment_web.gtpl")
		log.Println(t.Execute(w, nil))
	} else {
		//请求的是登录数据，那么执行登录的逻辑判断
		r.ParseForm()
		ss,_:=strconv.Atoi(r.Form["userid"][0])
		view.Addcomment(int32(ss),1,0,r.Form["content"][0])
	}
}
func main() {
	http.HandleFunc("/", showcomment) //设置访问的路由
	http.HandleFunc("/comment", comment_web)
	err := http.ListenAndServe(":1111", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}