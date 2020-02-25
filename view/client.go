package view

import (
	"SimpleComment/gen-go/comment"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)


func Addcomment(user_id int32, article_id int32, reply_id int32, content string)  (bool,error){

	client,err:=	runClient()
	if err != nil {
		panic(err)
		return false,err
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
		return false,err
	}
	println("success")
	if(res){
		return true,nil
	}else{
		return false,err
	}
}
func Getcomment(replyId int, articleId int) ([]*comment.Com,error) {

	client,err:=	runClient()
	if err != nil {
		panic(err)
		return nil,err
	}
	var r []*comment.Com
	r,err=client.Get(int32(replyId), int32(articleId))
	if err!=nil{
		panic(err)
		return nil,err
	}
	return r,nil
}
func DelteComment(id int)bool{
	client,err:=	runClient()
	if err != nil {
		panic(err)
		return false
	}
	fmt.Println(int32(id))

	r,err:=client.DeleteComment(int32(id))
	if r{
		return true
	}
		return false
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
