package main

import (
	"fmt"
	"net"
	"os"
)

// 2.接收服务端返回信息
func RecvMsg(conn net.Conn) {
	for {
		buf := make([]byte, 4*1024)
		n, err2 := conn.Read(buf)
		if err2 != nil {
			fmt.Println("conn.Read err = ", err2)
			return
		}

		fmt.Println(string(buf[:n]))
	}
}

// 3.发送消息给服务端
func SendMsg(conn net.Conn) {
	// 获取键盘输入数据
	for {
		buf := make([]byte, 4*1024)
		n, err := os.Stdin.Read(buf)
		if err != nil {
			fmt.Println("os.Stdin.Read err = ", err)
			return
		}

		//inputData := string(buf[:n])
		//inputData = strings.Replace(inputData, "\n", "", -1)

		// 发送信息给服务端
		conn.Write(buf[:n])
		//fmt.Println("发送给服务端的数据是: ", string(buf[:n]))
	}
}

func main() {
	// 1.连接服务器
	conn, err1 := net.Dial("tcp", "127.0.0.1:8005")
	if err1 != nil {
		fmt.Println("net.Dial err = ", err1)
		return
	}
	defer conn.Close()

	// 2.接收服务端返回信息
	go RecvMsg(conn)

	// 3.发送消息给服务端
	go SendMsg(conn)

	for {

	}
}
