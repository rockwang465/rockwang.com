package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

// 4.1创建struct，用于保存用户信息  -- 问题2:结构体要在外面定义，只要是其他函数有引用，这里必须定义到外面
type ClientInfo struct { // 创建结构体
	msg         chan string
	name        string
	userAddress string
}

// 4.2创建map字典和使用ClientInfo结构体，用于保存用户信息
var onlineUsers = make(map[string]ClientInfo) // 创建map字典

// 4.3创建一个channel，用于新增map值后，告知其他函数使用 -- 问题3:没有考虑到此处的妙用
var message = make(chan string) // 问题4: 这里写法必须要make进行初始化才行，否则没有开辟内存，后面会一直卡着的

// 9.1创建一个channel，用于判断用户是否退出
var isQuit = make(chan bool)

// 10.1创建一个channel，用于判断用户是否有正常的输入
var hasData = make(chan bool)

// 6.当message这个channel有内容时，则将内容放到每个Client的channel中
// 用于广播信息使用，当任何用户上线，或者发消息，都通过这个channel传输
func ManageCli() {
	for { // 问题6，最大问题在这，这里我没加for循环
		chatMsg := <-message
		for _, cli := range onlineUsers {
			cli.msg <- chatMsg
		}
	}
}

// 5.将上线信息发给当前客户端，告诉他们XXX上线了
func WriteMsgToClient(conn net.Conn, currentCli ClientInfo) {
	//userMsg := <-message // 5.1当读到4.4中的值后，才能进行后续操作

	//5.1这里时时读取ManageCli中写到当前客户端结构体中的信息，然后读取当前客户都的结构体中信息，返回给连接的客户端
	for respMsg := range currentCli.msg {
		//fmt.Println("回复给客户端的信息,告知谁上线了: ", respMsg)
		conn.Write([]byte(respMsg + "\n"))
	}
}

// 7.读取用户发来的消息
func ReadMsgForClient(conn net.Conn, currentCli ClientInfo) {
	buf := make([]byte, 4*1024) // 问题7:buf创建要放在for循环外面，否则client端程序退出再进就没反应了
	for {
		n, err := conn.Read(buf)
		//if err != nil {
		//	if err == io.EOF {
		//		continue
		//	} else if n == 0 {
		//		fmt.Println("conn.Read err = ", err, n)
		//		isQuit <- true
		//	}
		//}

		if n == 0 { // 说明用户断开连接了
			fmt.Println("n == 0, err = ", err)
			isQuit <- true // 9.2 当n=0，则说明用户退出了
			return
		}

		ReadMsg := string(buf[:n])
		ReadMsg = strings.Replace(ReadMsg, "\n", "", -1)
		fmt.Printf("#%v#\n",ReadMsg)
		listReadMsg := strings.Split(ReadMsg, "|")
		//fmt.Println("list Read Msg : ", listReadMsg)
		if ReadMsg == "who" && len(ReadMsg) == 3 { // 8.1 用户输入who查询当前在线用户
			conn.Write([]byte("User List :\n"))
			for _, v := range onlineUsers {
				whoOnline := "online username: " + v.name + ", address: " + v.userAddress + "\n"
				conn.Write([]byte(whoOnline))
			}
			// 8.2 重命名用户名功能 : 格式: rename|rock ; 长度:rename|r 最少长度为8 ; 包含rename
		} else if len(ReadMsg) >= 8 && "rename" == listReadMsg[0] {
			currentCli.name = listReadMsg[1]
			onlineUsers[currentCli.userAddress] = currentCli

		} else {
			message <- "[" + currentCli.name + "] : " + ReadMsg // 7.2 把读到的数据放入message， 而message会被自动写入到所有用户的msg中，并进行打印
		}

		hasData <- true // 10.2 到这里，肯定用户是有输入内容的，所以给他true表示当前有输入内容
	}
}

func HandUsers(conn net.Conn) {
	defer conn.Close()

	userAddr := conn.RemoteAddr().String() // 获取用户地址，这里一定要String()，否则字符类型就是*net.TCPAddr
	//fmt.Printf("userAddr 类型%T\n", userAddr) // 问题1: 不加.String() 时userAddr 类型*net.TCPAddr

	// 4.5 添加用户数据到结构体中
	currentCli := ClientInfo{make(chan string), userAddr, userAddr}
	onlineUsers[userAddr] = currentCli
	//fmt.Println(onlineUsers)

	// 4.4 当其他函数读到此channel有值，才会进行后续操作
	loginUserInfo := "user: [" + userAddr + "] login"
	message <- loginUserInfo // 4.4 当其他函数读到此channel有值，才会进行后续操作

	// 5.写上线信息给当前客户端，告诉他们XXX上线了
	go WriteMsgToClient(conn, currentCli) // 问题8: 这里在7.1的上面，否则客户端无任何反应

	// 7.1 告知当前用户，哪个程序是自己的
	currentCli.msg <- "I am here"

	// 7.读取用户发来的消息
	go ReadMsgForClient(conn, currentCli)

	for { // 加个for循环，让此协程永远不退出，除非手动退出 -- 问题5，这里我没想到

		// 9.超时处理（超过20秒无用户操作，则自动退出）
		select {
		case <-isQuit: // 9.3用户如果退出，则删除onlineUsers中对应的用户
			message <- "[" + onlineUsers[userAddr].name + "]" + " : logout" // 9.4 告知所有人此用户退出聊天室
			delete(onlineUsers, userAddr)
			fmt.Println(onlineUsers)
			return // 退出需要return
		case <-hasData: // 10.3 有数据，说明用户最近有输入，正常的，不走超时路线，不做任何操作
		case <-time.After(20 * time.Second): // 10.4 20秒超时
			message <- "[" + onlineUsers[userAddr].name + "]" + " : time out leave" // 10.5 超时退出通知
			delete(onlineUsers, userAddr)                                           // 超时则在map中删除用户
			return // 问题9: 老师这里正常，我这里一旦超时，则所有程序都退出了，暂时未找到原因，这里耽误了太久了，以后有能力再来解决bug吧
		}
	}

}

func main() {
	// 1.监听
	listener, err1 := net.Listen("tcp", "127.0.0.1:8005")
	if err1 != nil {
		fmt.Println("net.Listen err = ", err1)
		return
	}
	defer listener.Close()

	// 6.当message这个channel有内容时，则将内容放到每个Client的channel中
	// 用于广播信息使用，当任何用户上线，或者发消息，都通过这个channel传输
	go ManageCli()

	// 2.阻塞等待用户请求
	for { // for循环达到服务端永远不退出的效果
		conn, err2 := listener.Accept()
		if err2 != nil {
			fmt.Println("listener.Accept err = ", err2)
			return
		}
		//defer conn.Close()  // for循环里应该不需要conn.Close()了

		// 3.接收用户请求
		go HandUsers(conn)
	}
}
