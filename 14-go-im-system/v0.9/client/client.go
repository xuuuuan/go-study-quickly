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
	model      int
}

func NewClient(ip string, port int) *Client {
	client := Client{
		ServerIp:   ip,
		ServerPort: port,
		model:      999,
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

func (c *Client) menu() bool {
	var model int

	fmt.Println("Select one: 1. Public chat 2. Private chat 3. Update username 0. Exit")
	_, err := fmt.Scanln(&model)
	if err != nil {
		return false
	}
	if model >= 0 && model <= 3 {
		c.model = model
		return true
	} else {
		fmt.Println(">>>>Please enter a number within 0~3<<<<")
		return false
	}
}

func (c *Client) Run() {
	for c.model != 0 {
		for c.menu() == false {
			// 如果返回false就代表非法继续让用户选择
		}
		//根据不同的模式处理不同的业务
		switch c.model {
		case 1:
			// 公聊模式
			fmt.Println("selected public chat")
			break
		case 2:
			// 私聊模式
			fmt.Println("selected private chat")
			break
		case 3:
			// 更新用户名
			fmt.Println("selected update username")
			break
		}
	}
}

func main() {
	// 解析命令行参数
	flag.Parse()

	client := NewClient(serverIp, serverPort)
	if client == nil {
		fmt.Println(">>>>>>>>>>connect fail<<<<<<<<<<")
	}
	fmt.Println(">>>>>>>>>>connect success<<<<<<<<<<")

	client.Run()
}