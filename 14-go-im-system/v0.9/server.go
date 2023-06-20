package main

import (
	"fmt"
	"io"
	"net"
	"runtime"
	"sync"
	"time"
)

type Server struct {
	IP               string
	Port             int
	OnlineUserMap    map[string]*User
	mapLock          sync.RWMutex
	ServerMsgChannel chan string
}

func NewServer(ip string, port int) *Server {
	server := Server{
		IP:               ip,
		Port:             port,
		OnlineUserMap:    make(map[string]*User),
		ServerMsgChannel: make(chan string),
	}
	return &server
}

func (server *Server) listenServerMsgChannel() {
	for {
		msg := <-server.ServerMsgChannel
		// 将msg发送给所有在线用户
		server.mapLock.Lock()
		for _, user := range server.OnlineUserMap {
			user.sendMsg2Client(msg)
		}
		server.mapLock.Unlock()
	}
}

// Broadcast
// @Description: 广播消息
// @param user	 发送消息的用户
// @param msg	 发送的消息
func (server *Server) Broadcast(user *User, msg string) {
	sendMsg := time.Now().Format(time.DateTime) + " [" + user.Addr + "]" + user.Name + ":" + msg
	server.ServerMsgChannel <- sendMsg
}

// ConnHandler 处理当前连接业务
func (server *Server) ConnHandler(conn net.Conn) {
	// 1.新建一个用户
	user := NewUser(server, conn)
	// 2.用户上线
	user.online()

	// 判断用户是否在线
	userIsLive := make(chan bool)

	// 3.接收客户端发送的消息
	go func() {
		for {
			buff := make([]byte, 4096)
			n, err := conn.Read(buff)
			// control c 关闭一个客户端
			if n == 0 {
				user.offline()
				return
			}
			if err != nil && err != io.EOF {
				fmt.Println("Conn Read err:", err)
				return
			}
			// 提取用户消息(因为发送的时候会敲击回车,所以在这里我们要去除'\n')
			msg := string(buff[:n-1])
			user.doMessage(msg)

			// 用户在线
			userIsLive <- true
		}
	}()

	// 阻塞(因为是作为一个sub goroutine 不阻塞的话会结束)
	for {
		select {
		case <-userIsLive:
		case <-time.After(300 * time.Second):
			{
				user.sendMsg2Client("You've been forced offline.")
				time.Sleep(100 * time.Millisecond)
				err := conn.Close()
				if err != nil {
					fmt.Println("conn failed to close when forced offline.")
				}
				//return
				runtime.Goexit()
			}
		}
	}
}

func (server *Server) Start() {
	// 1.socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", server.IP, server.Port))
	if err != nil {
		fmt.Println("socket listen error:", err)
		return
	}
	fmt.Println("start listening...")

	// 5.socket close
	defer func(l net.Listener) {
		err := l.Close()
		if err != nil {
			fmt.Println("socket close error:", err)
			runtime.Goexit()
		}
	}(listener)

	// 2.启动监听ServerMsgChannel的goroutine
	go server.listenServerMsgChannel()

	for {
		// 3.accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener accept error:", err)
			continue
		}
		// 4.handle
		go server.ConnHandler(conn)
	}
}
