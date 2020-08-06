package main

import (
	"fmt"
	"io"
	"net"
)

type Client struct {
	C    chan string // 用户发送数据的管道
	Name string      // 用户名
	Addr string      // 网络地址
}

// 保存在线用户
var onlineMap map[string]Client

//var onlineMap = make(map[string]Client)

var message1 = make(chan string)

// 3.新开一个协程，接收消息，只要有消息来了，遍历map，给map每个成员都发送此消息
func Manager() {
	// 给map分配空间
	onlineMap = make(map[string]Client)

	for {
		msg := <-message1 // 没有消息前，这里被阻塞
		// 遍历map，给map每个成员都发送此消息
		for _, cli := range onlineMap {
			cli.C <- msg
		}

	}
}

// 4.4 新开一个协程，专门给当前客户端发送信息
func WriteMsgToClient1(cli Client, conn net.Conn) {
	for msg := range cli.C { // 给当前客户端发送信息
		conn.Write([]byte(msg + "\n"))
	}
}

// 4.6
func MakeMsg(cli Client, msg string) (buf string) {
	buf = "[" + cli.Addr + "]" + cli.Name + ": login"
	return
}

func HandleConn(conn net.Conn) { // 处理用户链接
	defer conn.Close()
	// 4.1 获取客户端的网络地址
	cliAddr := conn.RemoteAddr().String()

	// 4.2 创建一个结构体, 默认用户名和网络地址一样
	cli := Client{make(chan string), cliAddr, cliAddr}

	// 4.3 把结构体添加到map
	onlineMap[cliAddr] = cli

	// 4.4 新开一个协程，专门给当前客户端发送信息
	go WriteMsgToClient1(cli, conn)

	// 4.6 广播某个人在线
	//message1 <- "[" + cli.Addr + "]" + cli.Name + ": login"
	message1 <- MakeMsg(cli, "login")

	cli.C <- MakeMsg(cli, "I am here")

	// 5.接收用户发来的信息，并放入到用户map的chan中存放
	go func() {
		buf := make([]byte, 2048)
		for {
			n, err := conn.Read(buf)
			if err != nil {
				if err == io.EOF {
					continue
				} else {
					fmt.Println("conn.Read err = ", err)
				}
			}

			if n == 0 { //表示对方断开，或者出了问题
				fmt.Println("n==0, conn.Read err = ", err)
			}
			//msg := string(buf[:n-1])
			msg := string(buf[:n])

			//转发此内容（这样所有用户列表的chan中就会有此信息了）
			message1 <- MakeMsg(cli, msg)
		}

	}()

	for { // 4.5 特地不让程序结束的。如果结束了，当前HandleConn就关了

	}
}

func main() {
	// 1.监听
	listener, err := net.Listen("tcp", ":8007")
	if err != nil {
		fmt.Println("net.Listen err = ", err)
		return
	}

	defer listener.Close()

	// 3.新开一个协程，接收消息，只要有消息来了，遍历map，给map每个成员都发送此消息
	go Manager()

	// 2.主协程，循环阻塞等待用户链接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err = ", err)
			continue
		}

		go HandleConn(conn) // 处理用户链接

	}

}
