package main

import (
	"fmt"
	"net"
)

func main() {
	tcpAddr := "127.0.0.1:8001"
	// 1.主动连接服务器
	dia, err1 := net.Dial("tcp", tcpAddr)
	if err1 != nil {
		fmt.Println("net.Dial err = ", err1)
		return
	}

	defer dia.Close()

	// 2.发送数据
	n, err2 := dia.Write([]byte("hello"))
	if err2 != nil {
		fmt.Println("dia.Write err = ", err2)
		return
	}
	fmt.Println(n) // 5
}
