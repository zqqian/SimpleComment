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
	reply:=0
	if value, ok := r.Form["reply"];ok{
		reply,_=strconv.Atoi(value[0]) 	//获取回复id，默认为0全部显示
	}
	article:=0
	if value, ok := r.Form["article"];ok{
		article,_=strconv.Atoi(value[0]) //获取文章id，默认为0全部显示
	}
	res,err := view.Getcomment(reply,article)  //远程调用

	if err!=nil{
		panic(err)
	}
	for _,s:=range res { //显示评论
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
		rep,_=strconv.Atoi(value[0]) //获取回id
	}
	if r.Method == "GET" { //显示回复页面
		t, _ := template.ParseFiles("view/comment_web.gtpl")
		log.Println(t.Execute(w, rep))
	} else {
		ss,err:=strconv.Atoi(r.Form["userid"][0]) //获取用户id并转为int
		if err!=nil{
			panic(err)
		}
		r,err:=view.Addcomment(int32(ss),1,int32(rep),r.Form["content"][0])
		if err!=nil{
			panic(err)
		}
		http.Redirect(w,r,"/",301) //301跳转
	}
}
func delete_web(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		ss,err:=strconv.Atoi(r.Form["id"][0])
		if err!=nil{
			panic(err)
		}
		view.DelteComment(ss)
		http.Redirect(w,r,"/",301)
}
func main() {
	http.HandleFunc("/", showcomment) //显示评论
	http.HandleFunc("/comment", comment_web) //添加评论
	http.HandleFunc("/delete", delete_web) //删除评论
	err := http.ListenAndServe(":1111", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}