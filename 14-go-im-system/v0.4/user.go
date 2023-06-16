package main

import (
	"fmt"
	"net"
)

type User struct {
	Name        string
	Addr        string
	UserMsgChan chan string
	server      *Server
	connect     net.Conn
}

func NewUser(server *Server, conn net.Conn) *User {
	remoteAddr := conn.RemoteAddr().String()
	// 暂时先将user.Name设置为remoteAddr
	user := User{
		Name:        remoteAddr,
		Addr:        remoteAddr,
		UserMsgChan: make(chan string),
		server:      server,
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

func (user *User) online() {
	// 将用户信息加到OnlineUserMap中
	user.server.mapLock.Lock()
	user.server.OnlineUserMap[user.Name] = user
	user.server.mapLock.Unlock()
	// 广播该用户已经上线
	user.server.Broadcast(user, "is online.")
}

func (user *User) offline() {
	// 将用户信息加到OnlineUserMap中
	user.server.mapLock.Lock()
	delete(user.server.OnlineUserMap, user.Name)
	user.server.mapLock.Unlock()
	// 广播该用户已经下线
	user.server.Broadcast(user, "is offline.")
}

func (user *User) doBroadcastMsg(msg string) {
	user.server.Broadcast(user, msg)
}
