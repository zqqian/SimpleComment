package main

/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements. See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership. The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License. You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

import (
	"database/sql"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	_ "github.com/go-sql-driver/mysql"
	"os"
)
/*
type EchoServerImp struct {
}
func (e *EchoServerImp) Echo(req *echo.EchoReq) (*echo.EchoRes, error) {
	fmt.Printf("message from client: %v\n", req.GetMsg())
	res := &echo.EchoRes{
		Msg: req.GetMsg()+"aaaa",
	}
	return res, nil
}
*/
type CommentServer struct {

}

func (c *CommentServer) Add(name string, content string) (r bool, err error) {
	fmt.Println("add",name,"content ",content)
	db, err := sql.Open("mysql", "root:root@/simplecomment?charset=utf8")
	checkErr(err)
defer db.Close()
	//插入数据INSERT INTO `comment` (`id`, `username`, `context`, `time`) VALUES (NULL, 'un', 'cont', CURRENT_TIME());
	stmt, err := db.Prepare("INSERT INTO `comment` (`id`, `username`, `context`, `time`) VALUES (NULL, ?, ?, CURRENT_TIME())")
	checkErr(err)

	res, err := stmt.Exec(name, content)
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)
	return true,nil
}
func (c *CommentServer)Get() (r string, err error){


	db, err := sql.Open("mysql", "root:root@/simplecomment?charset=utf8")
	checkErr(err)
	defer db.Close()
	rows, err := db.Query("SELECT * FROM comment")
	checkErr(err)
	var s string
	s=""
	for rows.Next() {
		var id int
		var username string
		var context string
		var time string

		err = rows.Scan(&id, &username, &context,&time)
		checkErr(err)
		fmt.Println(id)
		s+=string(id)
		s+="\n"
	//	fmt.Println(username)
		s+=username
		s+="\n"
	//	fmt.Println(context)
		s+=context
		s+="\n"
	//	fmt.Println(time)
		s+=time
		s+="\n"
	}
	return s,nil
}

func main() {
	//flag.Usage = Usage
	//server := flag.Bool("server", false, "Run server")
	server := true
	//protocol := flag.String("P", "binary", "Specify the protocol (binary, compact, json, simplejson)")
	protocol:="compact"
	//framed := flag.Bool("framed", false, "Use framed transport")
	framed :=false
	//buffered := flag.Bool("buffered", false, "Use buffered transport")
	buffered :=true
	//addr := flag.String("addr", "localhost:9090", "Address to listen to")
	addr :="localhost:7777"
	//secure := flag.Bool("secure", false, "Use tls secure transport")
	secure :=false

		//flag.Parse()

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
		//Usage()
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

	if server {
		if err := runServer(transportFactory, protocolFactory, addr, secure); err != nil {
			fmt.Println("error running server:", err)
		}
	}
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}