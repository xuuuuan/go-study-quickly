package main

import (
	"fmt"
	"net"
	"strings"
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

// 向用户客户端发送消息
func (user *User) sendMsg2Client(msg string) {
	_, err := user.connect.Write([]byte(msg + "\n"))
	if err != nil {
		fmt.Printf("向%s客户端发送消息「%s」失败\n", user.Name, msg)
	}
}

// 用户上线
func (user *User) online() {
	// 将用户信息加到OnlineUserMap中
	user.server.mapLock.Lock()
	user.server.OnlineUserMap[user.Name] = user
	user.server.mapLock.Unlock()
	// 广播该用户已经上线
	user.server.Broadcast(user, "is online.")
}

// 用户下线
func (user *User) offline() {
	// 将用户信息加到OnlineUserMap中
	user.server.mapLock.Lock()
	delete(user.server.OnlineUserMap, user.Name)
	user.server.mapLock.Unlock()
	// 广播该用户已经下线
	user.server.Broadcast(user, "is offline.")
}

// 处理消息
func (user *User) doMessage(msg string) {
	// 在线用户列表
	if msg == "users" {
		user.server.mapLock.Lock()
		var userNameList []string
		for _, u := range user.server.OnlineUserMap {
			userNameList = append(userNameList, u.Name)
		}
		user.server.mapLock.Unlock()
		user.UserMsgChan <- "online users:" + strings.Join(userNameList, ",")
	} else if len(msg) > 7 && strings.HasPrefix(msg, "rename-") {
		newName := msg[7:]
		// 查看map里是否有该元素
		_, ok := user.server.OnlineUserMap[newName]
		if ok {
			user.UserMsgChan <- "The user name already exists."
		} else {
			user.server.mapLock.Lock()
			delete(user.server.OnlineUserMap, user.Name)
			user.Name = newName
			user.server.OnlineUserMap[newName] = user
			user.UserMsgChan <- "The user name is changed successfully."
			user.server.mapLock.Unlock()
		}
	} else {
		user.server.Broadcast(user, msg)
	}
}
