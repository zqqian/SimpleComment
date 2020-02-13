package main

import (
	"SimpleComment/gen-go/comment"
	"crypto/tls"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)


func addcomment(username ,content string)  bool{

	client,err:=	runClient()
	if err != nil {
		fmt.Println("failed. err: [%v]", err)
		return false
	}

	var res bool
	res,err=client.Add(username,content)

	if err != nil {
		fmt.Println(" failed. err: [%v]", err)
		return false
	}
	if(res){
		return true
	}else{
		return false
	}
}
func getcomment()  {

	client,err:=	runClient()
	if err != nil {
		fmt.Println("failed. err: [%v]", err)
		return
	}
	var commentlist string
	commentlist,err=client.Get()
	if err!=nil{
		fmt.Println("failed. err: [%v]", err)
		return
	}
	println(commentlist)
}
func runClient() (*comment.CommentClient,error) {
	transportFactory := thrift.NewTBufferedTransportFactory(8192)

	protocolFactory := thrift.NewTCompactProtocolFactory()
	addr:="localhost:7777"
	secure:=false
	var transport thrift.TTransport
	var err error
	if secure {
		cfg := new(tls.Config)
		cfg.InsecureSkipVerify = true
		transport, err = thrift.NewTSSLSocket(addr, cfg)
	} else {
		transport, err = thrift.NewTSocket(addr)
	}
	if err != nil {
		fmt.Println("Error opening socket:", err)
		return nil,err
	}
	transport = transportFactory.GetTransport(transport)
	if err := transport.Open(); err != nil {
		fmt.Println("Error opening socket:", err)
		return nil,err
	}
	return comment.NewCommentClientFactory(transport, protocolFactory),nil
}
func main() {

addcomment("testuser","tset a comment")

getcomment()

}