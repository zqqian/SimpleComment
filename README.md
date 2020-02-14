### Requirement

* Golang
* thrift 0.9.4
* [https://github.com/go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)

### **数据库**
在SimpleComment数据库下新建一张名为comment的表

字段：

| id   | int(11)   | 评论的id，从1开始自增   | 
|:----|:----|:----|
| username   | varchar(100)   | 评论的用户名   | 
| content   | text   | 评论的内容   | 
| time   | datetime   | 评论发表的时间，默认为CURRENT_TIME()   | 

### ### **文件列表**
SimpleComment

* gen-go  
* client.go //包括了一个客户端的简单实现
* comment.thrift  //thrift的IDL文件
* main.go   //服务端主程序


### thritft连接参数
addr :="localhost:7777" //默认端口7777

protocolFactory = thrift.NewTCompactProtocolFactory()

transportFactory = thrift.NewTBufferedTransportFactory(8192)

thrift IDL文件：comment.thrift

```
namespace go comment
struct com{
1:i32 id,
2:string username,
3:string content,
4:string time,
}
service Comment {
    bool add(1: string name,2:string content),
    list<com> get()
}
```
定义了两个服务，分别为添加评论和获取评论列表
### 添加评论
通过调用这个函数可以往数据库中新添加一条评论，需要提供评论的用户名和评论内容。评论的时间默认为当前的时间。

函数定义：

>func (c *CommentServer) Add(name string, content string) (r bool, err error)

请求参数

| 字段   | 说明   | 类型   | 是否必填   | 
|:----|:----|:----|:----|
| name   | 评论的用户名   | string   | 是   | 
| content   | 评论的内容   | string    | 是   | 

返回参数

| 字段   | 说明   | 类型   | 备注   | 
|:----|:----|:----|:----|
| r   | 是否成功   | bool   | 成功为true，失败为false   | 
| err   | 错误信息   | error   |    | 

示例：

```
res,err=client.Add("test-user","test-comment")
```
### 获取评论列表
通过这个函数来获得当前的评论列表。

返回值是一个结构体切片，

结构体中定义了四个类型，分别是

评论id 

评论用户名

评论内容

评论时间

函数定义：

>func (c *CommentServer)Get() (r []*comment.Com, err error)

请求参数：无

返回参数

| 字段 | 说明 | 类型 | 备注 | 
|:----:|:----:|:----:|:----|:----:|
| r | 评论列表 | []*comment.Com | 其中comment.Com定义为：type Com struct {   Id       int32     Username string    Content  string    Time     string   }  默认按照时间顺序排列   | 
| err | 错误信息 | error |    | 

示例：

```
var r []*comment.Com
r,err=client.Get()
 if err!=nil{
   panic(err)
}
for _,s:=range r {
   println(s.Username)
   println(s.Content)
   println(s.Id)
   println(s.Time)
}
```















