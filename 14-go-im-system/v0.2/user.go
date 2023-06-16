package main

import (
	"fmt"
	"net"
)

type User struct {
	Name        string
	Addr        string
	UserMsgChan chan string
	connect     net.Conn
}

func NewUser(conn net.Conn) *User {
	remoteAddr := conn.RemoteAddr().String()
	// 暂时先将user.Name设置为remoteAddr
	user := User{
		Name:        remoteAddr,
		Addr:        remoteAddr,
		UserMsgChan: make(chan string),
		connect:     conn,
	}
	// 起一个goroutine
	go user.listenUserMsgChan()
	return &user
}

// 监听UserMsgChan有msg就向client发送
func (user *User) listenUserMsgChan() {
	for {
		msg := <-user.UserMsgChan
		_, err := user.connect.Write([]byte(msg + "\n"))
		if err != nil {
			fmt.Printf("向%s发送消息「%s」失败\n", user.Name, msg)
			continue
		}
	}
}
