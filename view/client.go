package view

import (
	"SimpleComment/gen-go/comment"
	"git.apache.org/thrift.git/lib/go/thrift"
)


func Addcomment(user_id int32, article_id int32, reply_id int32, content string)  bool{

	client,err:=	runClient()
	if err != nil {
		panic(err)
		return false
	}

	var res bool
	cc:=new(comment.Com)
	cc.Content=content
	cc.ReplyId=reply_id
	cc.ArticleId=article_id
	cc.UserId=user_id

	res,err=client.AddComment(cc)

	if err != nil {
		panic(err)
		return false
	}
	println("success")
	if(res){
		return true
	}else{
		return false
	}
}
func Getcomment() ([]*comment.Com,error) {

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
	var transport thrift.TTransport
	var err error
	transport, err = thrift.NewTSocket(addr)
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
//func main() {
//
//	Addcomment(1,1,0,"111")
//
// 	r,err :=Getcomment()
//
//	 if err!=nil{
//		panic(err)
//	}
//	for _,s:=range r {
//		println(s.Username)
//		println(s.Content)
//		println(s.Id)
//		println(s.Time)
//	}
//
//}