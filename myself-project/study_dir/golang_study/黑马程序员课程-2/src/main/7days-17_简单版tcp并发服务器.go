package main

import (
	"fmt"
	"io"
	"net"
	"strings"
)

func HandInfo(conn net.Conn) {
	// 3.获取用户地址信息
	userAddr := conn.RemoteAddr().String()
	fmt.Printf("[%v]: connected successful\n", userAddr)

	// 4.读取用户发来的数据
	for {
		buf := make([]byte, 4*1024)
		n, err1 := conn.Read(buf)
		if err1 != nil {
			if err1 == io.EOF { // 当用户输入exit时，退出，则为空
				fmt.Printf("[%v]: bye bye\n", userAddr)
				break
			} else {
				fmt.Println("HandInfo conn.Read err = ", err1)
				return
			}
		}
		fmt.Printf("[%v]: %v\n", userAddr, string(buf[:n]))
		userData := string(buf[:n]) // 这里可能存在着换行，导致下面if判断无法进行

		//compareData := strings.Replace(userData, " ", "", -1)
		//compareData = strings.Replace(compareData, "\t", "", -1)
		compareData := strings.Replace(userData, "\n", "", -1) // 测试发现多了一个"\n"

		if compareData == "exit" || compareData == "EXIT" {
			fmt.Printf("[%v]: bye-bye ! \n", userAddr)
			//fmt.Println(len(compareData))  // 通过查看长度也可以检测出来
			//fmt.Printf("[%v]", userData)  // 测试发现多了一个"\n"，所以上面值替换了"\n"
			return
		}
		// 5.回复用户大写内容
		dataUpper := []byte(strings.ToUpper(userData))
		conn.Write(dataUpper)
	}
}

func main() {
	// 1. 监听
	listener, err1 := net.Listen("tcp", "127.0.0.1:8002")
	if err1 != nil {
		fmt.Println("net.Listen err = ", err1)
		return
	}
	defer listener.Close()

	// 2.阻塞，接收多个用户
	for {
		conn, err2 := listener.Accept()
		if err2 != nil {
			fmt.Println("conn.Accept err = ", err2)
			return
		}

		// 3.处理用户请求，新建一个协程
		go HandInfo(conn)
	}
}

// Rock总结:
// 1. 容易忘记close
// 2. 获取的内容，可能存在\t \n 空格等，需要替换掉。
