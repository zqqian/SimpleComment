package main

import (
	"SimpleComment/gen-go/comment"
	"crypto/tls"
	"git.apache.org/thrift.git/lib/go/thrift"
)


func addcomment(username ,content string)  bool{

	client,err:=	runClient()
	if err != nil {
		panic(err)
		return false
	}

	var res bool
	res,err=client.Add(username,content)

	if err != nil {
		panic(err)
		return false
	}
	if(res){
		return true
	}else{
		return false
	}
}
func getcomment() ([]*comment.Com,error) {

	client,err:=	runClient()
	if err != nil {
		panic(err)
		return nil,err
	}
	var r []*comment.Com
	r,err=client.Get()
	if err!=nil{
		panic(err)
		return nil,err
	}
	return r,nil
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
		panic(err)
		return nil,err
	}
	transport = transportFactory.GetTransport(transport)
	if err := transport.Open(); err != nil {
		panic(err)
		return nil,err
	}
	return comment.NewCommentClientFactory(transport, protocolFactory),nil
}
func main() {

	addcomment("testuser","tset a comment")

 	r,err :=getcomment()

	 if err!=nil{
		panic(err)
	}
	for _,s:=range r {
		println(s.Username)
		println(s.Content)
		println(s.Id)
		println(s.Time)
	}

}