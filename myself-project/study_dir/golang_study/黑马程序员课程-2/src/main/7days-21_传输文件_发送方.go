package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func sendFile(path string, conn net.Conn) {
	// A.打开文件
	fp, err1 := os.Open(path)
	if err1 != nil {
		fmt.Println("os.Open err = ", err1)
	}
	defer fp.Close()

	// B.读取文件
	buf := make([]byte, 1024)
	for {
		n, err2 := fp.Read(buf)
		if err2 != nil {
			if err2 == io.EOF {
				fmt.Println("文件发送完毕")
				return
			} else {
				fmt.Println("fp.Read err ", err2)
			}
		}

		// C.写文件，发送给服务端
		_, err3 := conn.Write(buf[:n])
		if err3 != nil {
			fmt.Println("conn.Write err = ", err3)
		}
	}
}

func main() {

	// 1.用户输入需要发送的文件路径
	fmt.Println("请输入需要传输的文件名:")
	var path string
	fmt.Scan(&path)
	fmt.Println("path = ", path)

	// 2.获取文件名
	fileInfo, err2 := os.Stat(path)
	if err2 != nil {
		fmt.Println("os.Stat err = ", err2)
	}
	//fmt.Println(fileInfo.Name())  // 文件名

	// 3.连接服务器
	conn, err1 := net.Dial("tcp", "127.0.0.1:8004")
	if err1 != nil {
		fmt.Println("net.Dial err = ", err1)
	}

	defer conn.Close()

	// 4.发送文件名给服务端
	_, err3 := conn.Write([]byte(fileInfo.Name()))
	if err3 != nil {
		fmt.Println("conn.Write err = ", err3)
	}

	// 5.读取服务器返回信息
	buf := make([]byte, 1024*4)
	resp, err4 := conn.Read(buf)
	if err4 != nil {
		fmt.Println("conn.Read err = ", err4)
		return
	}

	if string(buf[:resp]) == "ok" {
		sendFile(path, conn)
	}
}
