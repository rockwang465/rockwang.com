package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	// 1.监听
	listenr, err1 := net.Listen("tcp", "127.0.0.1:8001")
	if err1 != nil {
		if err1 == io.EOF {
			fmt.Println("ending")
			return
		} else {
			fmt.Println("net.Listen err = ", err1)
			return
		}
	}

	defer listenr.Close()

	// 2.阻塞，等待用户连接
	conn, err2 := listenr.Accept()
	if err2 != nil {
		fmt.Println("listenr.Accept err = ", err2)
		return
	}
	defer conn.Close() // 关闭当前连接

	// 3.接收用户请求
	buf := make([]byte, 4*1024)
	num, err3 := conn.Read(buf)
	if err3 != nil {
		fmt.Println("conn.Read err = ", err3)
		return
	}

	fmt.Println("buf = ", string(buf[:num]))
}
