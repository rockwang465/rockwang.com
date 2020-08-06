package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func returnData(conn net.Conn) {
	// 4.接收服务器返回内容
	buf2 := make([]byte, 4*1024)
	res1, err4 := conn.Read(buf2)
	if err4 != nil {
		fmt.Println("conn.Read err = ", err4)
		return
	}
	fmt.Println(string(buf2[:res1]))
}

func main() {
	// 1.连接服务器
	conn, err1 := net.Dial("tcp", "127.0.0.1:8002")
	if err1 != nil {
		fmt.Println("net.Listen err = ", err1)
		return
	}

	defer conn.Close()

	//go func() {
	//	// 4.接收服务器返回内容
	//	buf2 := make([]byte, 4*1024)
	//	res1, err4 := conn.Read(buf2)
	//	if err4 != nil {
	//		fmt.Println("conn.Read err = ", err4)
	//		return
	//	}
	//	fmt.Println(string(buf2[:res1]))
	//}()

	for {
		// 2.键盘输入聊天内容
		buf := make([]byte, 4*1024)
		n, err2 := os.Stdin.Read(buf)
		if err2 != nil {
			fmt.Println("os.Stdin.Read err = ", err2)
			return
		}

		//fmt.Println(string(buf[:n]))
		inputData := string(buf[:n])
		inputData = strings.Replace(inputData, "\n", "", -1)

		if inputData == "exit" || inputData == "exit" {
			return
		}

		// 3.发送内容给服务端
		_, err3 := conn.Write(buf[:n])
		if err3 != nil {
			fmt.Println("conn.Write err = ", err3)
			return
		}

		returnData(conn)
	}
}
