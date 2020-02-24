package main

import (
	"SimpleComment/controllers"
	"SimpleComment/gen-go/comment"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	_ "github.com/go-sql-driver/mysql"
)



func main() {
	if err := runServer(); err != nil {
		fmt.Println("error running server:", err)
	}
}

func runServer() error {
	addr :="localhost:7777"
	var protocolFactory thrift.TProtocolFactory
	protocolFactory = thrift.NewTCompactProtocolFactory()
	var transportFactory thrift.TTransportFactory
	transportFactory = thrift.NewTBufferedTransportFactory(8192)
	var transport thrift.TServerTransport
	var err error
	transport, err = thrift.NewTServerSocket(addr)
	if err != nil {
		return err
	}
	fmt.Printf("%T\n", transport)
	processor  :=comment.NewCommentProcessor(&controllers.CommentServer{})
	server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)
	fmt.Println("Starting the simple server... on ", addr)
	return server.Serve()
}
