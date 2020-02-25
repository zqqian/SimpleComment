package controllers

import (
	"SimpleComment/gen-go/comment"
	"SimpleComment/models"
	"errors"
	"fmt"
	"regexp"
	"strconv"
)
type CommentServer struct {
}
func FilteredSQLInject(toMatchStr string) bool {//SQL注入检查
	str := `(?:')|(?:--)|(/\\*(?:.|[\\n\\r])*?\\*/)|(\b(select|update|and|or|delete|insert|trancate|char|chr|into|substr|ascii|declare|exec|count|master|into|drop|execute)\b)`
	re, err := regexp.Compile(str)
	if err != nil {
		panic(err.Error())
		return false
	}
	return re.MatchString(toMatchStr)}
func checkUserExist(userId int32) bool {//添加评论之前检查用户是否存在
	sql:="SELECT * FROM `user` WHERE `userid` = "+strconv.Itoa(int(userId))
	if models.SqlExe(sql){
		return true
	}
	return false
}
func checkComment(cc *comment.Com) bool {//检查评论内容是否有效
	if cc.Content==""||cc.UserId==0||cc.ArticleId==0{
		return false
	}
	if !checkUserExist(cc.UserId){
		return false
	}
	if FilteredSQLInject(cc.Content){
		return false
	}
	return true
}
func checkCommentExist(commentId int32) bool {//删除评论之前检查评论状态
	sql:="SELECT * FROM `comment` WHERE `id` = "+string(commentId)
	if models.SqlExe(sql){
		return true
	}
	return false
}
func (c *CommentServer) AddComment(cc *comment.Com) (r bool, err error) {//添加一条评论
	//if !checkComment(cc){
	//	return false,errors.New("user id not exist")
	//}
	sql:="INSERT INTO `comment` (`id`, `userid`, `article_id`, `reply_id`, `context`, `time`) VALUES (NULL, '"+strconv.Itoa(int(cc.UserId))+"', '"+strconv.Itoa(int(cc.ArticleId))+"', '"+strconv.Itoa(int(cc.ReplyId))+"', '"+cc.Content+"', CURRENT_TIME());"
	if models.SqlExe(sql){
		return true,nil
	}else{
		return false,err
	}
}
func (c *CommentServer)Get(replyId int32) (r []*comment.Com, err error){//获取评论列表
	db:=models.Sqlconn()
	defer db.Close()
	sql:="SELECT id,comment.userid,user.username,context,time FROM `comment`,`user` WHERE comment.userid=user.userid and reply_id="+strconv.Itoa(int(replyId))
	fmt.Println(sql)
	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
		return nil,err

	}
	co:= []*comment.Com{}
	for rows.Next() {
		var comm comment.Com
		err = rows.Scan(&comm.Id,&comm.UserId,&comm.Username, &comm.Content, &comm.Time)
		if err != nil {
			panic(err)
			return nil,errors.New("fail to load comment")
		}
		co=append(co,&comm)
	}
	return co,nil
}
func (c *CommentServer)DeleteComment(commentId int32) (r bool, err error){//删除评论
//	if !checkCommentExist(commentId){
//	return false,errors.New("comment id not exist")
//}
	println(strconv.Itoa(int(commentId)))
	sql:="DELETE FROM `comment` WHERE `comment`.`id` = "+strconv.Itoa(int(commentId))
	if models.SqlExe(sql){
		return true,nil
	}else{
		return false,err
	}
}