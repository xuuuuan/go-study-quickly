package main

import (
	"fmt"
	"net"
	"runtime"
)

type Server struct {
	IP   string
	Port int
}

func NewServer(ip string, port int) *Server {
	server := Server{
		IP:   ip,
		Port: port,
	}
	return &server
}

func (server *Server) ConnHandler(conn net.Conn) {
	// 处理当前连接业务
	fmt.Println("connect success ", conn.RemoteAddr())
}

func (server *Server) Start() {
	// 1.socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", server.IP, server.Port))
	if err != nil {
		fmt.Println("socket listen error:", err)
		return
	}
	fmt.Println("start listening...")

	// 4.socket close
	defer func(l net.Listener) {
		err := l.Close()
		if err != nil {
			fmt.Println("socket close error:", err)
			runtime.Goexit()
		}
	}(listener)

	for {
		// 2.accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener accept error:", err)
			continue
		}
		// 3.handle
		go server.ConnHandler(conn)
	}
}
