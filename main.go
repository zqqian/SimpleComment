package main


import (
	"SimpleComment/gen-go/comment"
	"database/sql"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

type CommentServer struct {
}

func (c *CommentServer) Add(name string, content string) (r bool, err error) {//添加一条评论
	fmt.Println("add",name,"content ",content)
	db, err := sql.Open("mysql", "root:root@/simplecomment?charset=utf8")
	checkErr(err)
	defer db.Close()
	stmt, err := db.Prepare("INSERT INTO `comment` (`id`, `username`, `content`, `time`) VALUES (NULL, ?, ?, CURRENT_TIME())")
	checkErr(err)

	res, err := stmt.Exec(name, content)
	checkErr(err)

	id ,err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)
	return true,nil
}
func (c *CommentServer)Get() (r []*comment.Com, err error){//获取评论列表
	db, err := sql.Open("mysql", "root:root@/simplecomment?charset=utf8")
	checkErr(err)

	defer db.Close()

	rows, err := db.Query("SELECT * FROM comment")
	checkErr(err)
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
		checkErr(err)
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
	protocol:="compact"
	framed :=false
	buffered :=true
	addr :="localhost:7777"
	secure :=false

	var protocolFactory thrift.TProtocolFactory
	switch protocol {
	case "compact":
		protocolFactory = thrift.NewTCompactProtocolFactory()
	case "simplejson":
		protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
	case "json":
		protocolFactory = thrift.NewTJSONProtocolFactory()
	case "binary", "":
		protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
	default:
		fmt.Fprint(os.Stderr, "Invalid protocol specified", protocol, "\n")
		os.Exit(1)
	}
	var transportFactory thrift.TTransportFactory
	if buffered {
		transportFactory = thrift.NewTBufferedTransportFactory(8192)
	} else {
		transportFactory = thrift.NewTTransportFactory()
	}
	if framed {
		transportFactory = thrift.NewTFramedTransportFactory(transportFactory)
	}


		if err := runServer(transportFactory, protocolFactory, addr, secure); err != nil {
			fmt.Println("error running server:", err)
		}

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}