package main

import (
	"flag"
	"fmt"
	"net"
)

var serverIp string
var serverPort int

type Client struct {
	ServerIp   string
	ServerPort int
	Name       string
	connect    net.Conn
}

func NewClient(ip string, port int) *Client {
	client := Client{
		ServerIp:   ip,
		ServerPort: port,
	}

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", ip, port))
	if err != nil {
		fmt.Println("net.Dial err:", err)
		return nil
	}
	client.connect = conn
	return &client
}

func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "set server ip")
	flag.IntVar(&serverPort, "port", 6789, "set server port")
	/*
		效果如下 在命令行输入-h的提示
		./client -h
		Usage of ./client:
		-ip string
			set server ip (default "127.0.0.1")
		-port int
			set server port (default 6789)
	*/
}

func main() {
	// 解析命令行参数
	flag.Parse()

	client := NewClient(serverIp, serverPort)
	if client == nil {
		fmt.Println(">>>>>>>>>>connect fail<<<<<<<<<<")
	}
	fmt.Println(">>>>>>>>>>connect success<<<<<<<<<<")

	select {}
}
