package main


//import (
//	"fmt"
//	"net"
//	"time"
//)

//func main() {
//
//}
//
//package main
//
//import (
//"fmt"
//"io"
//"net"
//"strings"
//)
//
//// 4.1创建struct，用于保存用户信息  -- 问题2:结构体要在外面定义，只要是其他函数有引用，这里必须定义到外面
//type ClientInfo struct { // 创建结构体
//	msg         chan string
//	name        string
//	userAddress string
//}
//
//// 4.2创建map字典和使用ClientInfo结构体，用于保存用户信息
//var onlineUsers = make(map[string]ClientInfo) // 创建map字典
//
//// 4.3创建一个channel，用于新增map值后，告知其他函数使用 -- 问题3:没有考虑到此处的妙用
//var message = make(chan string) // 问题4: 这里写法必须要make进行初始化才行，否则没有开辟内存，后面会一直卡着的
//
//// 6.当message这个channel有内容时，则将内容放到每个Client的channel中
//// 用于广播信息使用，当任何用户上线，或者发消息，都通过这个channel传输
//func ManageCli() {
//	for { // 问题6，最大问题在这，这里我没加for循环
//		chatMsg := <-message
//		for _, cli := range onlineUsers {
//			cli.msg <- chatMsg
//		}
//	}
//}
//
//// 5.将上线信息发给当前客户端，告诉他们XXX上线了
//func WriteMsgToClient(conn net.Conn, currentCli ClientInfo) {
//	//userMsg := <-message // 5.1当读到4.4中的值后，才能进行后续操作
//
//	//5.1这里时时读取ManageCli中写到当前客户端结构体中的信息，然后读取当前客户都的结构体中信息，返回给连接的客户端
//	for respMsg := range currentCli.msg {
//		//fmt.Println("回复给客户端的信息,告知谁上线了: ", respMsg)
//		conn.Write([]byte(respMsg + "\n"))
//	}
//}
//
//// 7.读取用户发来的消息
//func ReadMsgForClient(conn net.Conn) {
//	buf := make([]byte, 4*1024) // 问题7:buf创建要放在for循环外面，否则client端程序退出再进就没反应了
//	for {
//		n, err := conn.Read(buf)
//		if err != nil {
//			if err == io.EOF {
//				continue
//			} else {
//				fmt.Println("conn.Read err = ", err)
//			}
//		}
//
//		if n == 0 {
//			fmt.Println("n == 0, err = ", err)
//		}
//
//		ReadMsg := string(buf[:n])
//		ReadMsg = strings.Replace(ReadMsg, "\n", "", -1)
//		//fmt.Printf("###%v###\n", ReadMsg)
//		if ReadMsg == "who" && len(ReadMsg) == 3 {
//			conn.Write([]byte("User List :"))
//			for _, v := range onlineUsers {
//				//fmt.Printf("online username: %v , address: %v ",v.name,v.userAddress )
//				whoOnline := "online username: " + v.name + ", address: " + v.userAddress + "\n"
//				conn.Write([]byte(whoOnline))
//			}
//		} else {
//			message <- ReadMsg // 把读到的数据放入message， 而message会被自动写入到所有用户的msg中，并进行打印
//		}
//	}
//}
//
//func HandUsers(conn net.Conn) {
//	defer conn.Close()
//
//	userAddr := conn.RemoteAddr().String() // 获取用户地址，这里一定要String()，否则字符类型就是*net.TCPAddr
//	//fmt.Println(userAddr)                   // 127.0.0.1:51487
//	//fmt.Printf("userAddr 类型%T\n", userAddr) // 问题1: 不加.String() 时userAddr 类型*net.TCPAddr
//	//fmt.Printf("用户 [%v] 上线了\n", userAddr)
//
//	// 4.5 添加用户数据到结构体中
//	currentCli := ClientInfo{make(chan string), userAddr, userAddr}
//	onlineUsers[userAddr] = currentCli
//	//fmt.Println(onlineUsers)
//
//	//message <- "true" // 4.4 当其他函数读到此channel有值，才会进行后续操作
//	loginUserInfo := "user: [" + userAddr + "] login"
//	message <- loginUserInfo // 4.4 当其他函数读到此channel有值，才会进行后续操作
//
//	// 5.写上线信息给当前客户端，告诉他们XXX上线了
//	go WriteMsgToClient(conn, currentCli)
//
//	// 7.1 告知当前用户，哪个程序是自己的
//	currentCli.msg <- "I am here"
//
//	// 7.读取用户发来的消息
//	go ReadMsgForClient(conn)
//
//	for { // 加个for循环，让此协程永远不退出，除非手动退出 -- 问题5，这里我没想到
//	}
//
//}
//
//func main() {
//	// 1.监听
//	listener, err1 := net.Listen("tcp", "127.0.0.1:8005")
//	if err1 != nil {
//		fmt.Println("net.Listen err = ", err1)
//		return
//	}
//	defer listener.Close()
//
//	// 6.当message这个channel有内容时，则将内容放到每个Client的channel中
//	// 用于广播信息使用，当任何用户上线，或者发消息，都通过这个channel传输
//	go ManageCli()
//
//	// 2.阻塞等待用户请求
//	for { // for循环达到服务端永远不退出的效果
//		conn, err2 := listener.Accept()
//		if err2 != nil {
//			fmt.Println("listener.Accept err = ", err2)
//			return
//		}
//		//defer conn.Close()  // for循环里应该不需要conn.Close()了
//
//		// 3.接收用户请求
//		go HandUsers(conn)
//	}
//}
