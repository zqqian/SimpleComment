package main

import (
	"SimpleComment/gen-go/comment"
	"database/sql"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	_ "github.com/go-sql-driver/mysql"
)

type CommentServer struct {
}

func (c *CommentServer) Add(name string, content string) (r bool, err error) {//添加一条评论
	db, err := sql.Open("mysql", "root:root@/simplecomment?charset=utf8")
	if err != nil {
		panic(err)
		return false,err
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO `comment` (`id`, `username`, `content`, `time`) VALUES (NULL, ?, ?, CURRENT_TIME())")
	if err != nil {
		panic(err)
		return false,err

	}

	res, err := stmt.Exec(name, content)
	if err != nil {
		panic(err)
		return false,err

	}

	id ,err := res.LastInsertId()
	if err != nil {
		panic(err)
		return false,err

	}
	fmt.Println(id)

	return true,nil
}
func (c *CommentServer)Get() (r []*comment.Com, err error){//获取评论列表
	db, err := sql.Open("mysql", "root:root@/simplecomment?charset=utf8")
	if err != nil {
		panic(err)
		return nil,err
	}

	defer db.Close()

	rows, err := db.Query("SELECT * FROM comment")
	if err != nil {
		panic(err)
		return nil,err

	}
	type com	struct {
		id int
		username string
		content string
		time string
	}
 	co:= []*comment.Com{}
	for rows.Next() {
		var id int32
		var username string
		var content string
		var time string

		err = rows.Scan(&id, &username, &content, &time)
		if err != nil {
			panic(err)
			return nil,err

		}
		fmt.Println(id)
		var comm comment.Com
		comm.Username = username
		comm.Content = content
		comm.Id = id
		comm.Time = time
		co=append(co,&comm)
	}

	return co,nil
}

func main() {

	addr :="localhost:7777"

	var protocolFactory thrift.TProtocolFactory
	protocolFactory = thrift.NewTCompactProtocolFactory()

	var transportFactory thrift.TTransportFactory
	transportFactory = thrift.NewTBufferedTransportFactory(8192)

	if err := runServer(transportFactory, protocolFactory, addr, false); err != nil {
		fmt.Println("error running server:", err)
	}

}

