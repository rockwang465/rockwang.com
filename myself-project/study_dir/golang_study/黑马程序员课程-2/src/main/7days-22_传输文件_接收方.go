package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

func recvFile(fileName string, conn net.Conn) {
	// A.创建文件
	file, err1 := os.Create(fileName)
	if err1 != nil {
		fmt.Println("os.Create err = ", err1)
		return
	}

	// B.读取用户发来的数据
	buf := make([]byte, 1024*4)
	for {
		n, err2 := conn.Read(buf)
		if err2 != nil {
			if err2 == io.EOF {
				fmt.Printf("[%v]文件接收完毕\n", fileName)
				time.Sleep(1 * time.Second)
				//return
			} else {
				fmt.Println("conn.Read err = ", err2)
				return
			}

			if n == 0{ // 必须在上面或这里加一个return/break的操作，否则死循环了
				fmt.Println(" n == 0 ,文件也接收完毕了")
				break
			}
		}
		// C.写入到文件中
		_,err3 := file.Write(buf[:n])
		if err3 != nil {
			fmt.Println("file.Write err = ", err3)
			return
		}
	}
}

func main() {
	// 1.监听
	listener, err1 := net.Listen("tcp", "127.0.0.1:8004")
	if err1 != nil {
		fmt.Println("net.Listen err = ", err1)
		return
	}
	defer listener.Close()

	// 2.阻塞等待用户连接
	conn, err2 := listener.Accept()
	if err2 != nil {
		fmt.Println("listener.Accept err = ", err2)
		return
	}
	defer conn.Close()

	// 3.读取客户端发来的文件名
	buf := make([]byte, 1024)
	n, err3 := conn.Read(buf)
	if err3 != nil {
		fmt.Println("conn.Read err = ", err3)
		return
	}

	// 4.回复接收到了文件名
	_, err4 := conn.Write([]byte("ok"))
	if err4 != nil {
		fmt.Println("conn.Write err = ", err4)
		return
	}
	fileName := string(buf[:n])

	recvFile(fileName, conn)
}
