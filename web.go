package main

import (
	"SimpleComment/view"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func showcomment(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w,"<html>\r\n<head>\r\n<title></title>\r\n</head>\r\n<body>")
	r.ParseForm()
	val:=0
	if value, ok := r.Form["reply"];ok{
		val,_=strconv.Atoi(value[0])
	}
	res,err := view.Getcomment(val)

	if err!=nil{
		panic(err)
	}
	for _,s:=range res {
		fmt.Fprintln(w, "用户名："+s.Username)
		fmt.Fprintln(w, "评论内容："+s.Content)
		fmt.Fprintln(w, "发表时间："+s.Time)
		fmt.Fprintln(w,"<a href=/delete?id="+strconv.Itoa(int(s.Id))+">删除</a>" )
		if val==0{
			fmt.Fprintln(w,"<a href=/?reply="+strconv.Itoa(int(s.Id))+">查看回复</a></br>" )
		}
	}
	fmt.Fprintln(w,"</br><a href=/comment?reply="+strconv.Itoa(val)+">添加评论</a></br>" )

}

func comment_web(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	rep:=0
	if value, ok := r.Form["reply"];ok{
		rep,_=strconv.Atoi(value[0])
	}
	if r.Method == "GET" {
		t, _ := template.ParseFiles("view/comment_web.gtpl")
		log.Println(t.Execute(w, rep))
	} else {
		ss,_:=strconv.Atoi(r.Form["userid"][0])
		view.Addcomment(int32(ss),1,int32(rep),r.Form["content"][0])
		http.Redirect(w,r,"/",301)
	}
}
func delete_web(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		ss,_:=strconv.Atoi(r.Form["id"][0])
		view.DelteComment(ss)
		http.Redirect(w,r,"/",301)
}
func main() {
	http.HandleFunc("/", showcomment) //设置访问的路由
	http.HandleFunc("/comment", comment_web)
	http.HandleFunc("/delete", delete_web)
	err := http.ListenAndServe(":1111", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}