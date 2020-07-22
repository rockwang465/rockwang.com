package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// 1.检测传参个数
	list := os.Args
	fmt.Println(list)

	//fmt.Println(len(list))  // 3
	if len(list) != 3 {
		fmt.Println("Error : 请输入2个文件名")
		fmt.Printf("Usage : %v filename1 filename2 \n" , list[0])
		return
	}

	// 2.比对传参文件名是否相同
	if list[1] == list[2] {
		fmt.Println("Error : 请输入两个不同的文件名！")
		os.Exit(2)
	}

	// 3.打开源文件，创建新文件
	src_file_name := list[1]
	dsc_file_name := list[2]
	src1, err1 := os.Open(src_file_name)
	if err1 != nil {
		fmt.Println("Error : os.Open = ", err1)
		return
	}

	dsc1, err2 := os.Create(dsc_file_name)
	if err2 != nil {
		fmt.Println("Error : os.Create = ", err2)
		return
	}

	defer src1.Close()  // 最容易忘记的地方，就是close掉文件
	defer dsc1.Close()

	// 4.读源文件，写入新文件
	buf := make([]byte, 4*1024)
	for {
		n, err3 := src1.Read(buf)  // res为读取的数量，为int数字；读取的内容放在了buf中
		if err3 != nil {
			if err3 == io.EOF {
				fmt.Println("Info : src1.Read = ", err3)
				break
			} else {
				fmt.Println("Error : src1.Read = ", err3)
			}
		}
		//fmt.Println(string(buf[:n]))  // 打印确认正常输出了

		_, err4 := dsc1.Write(buf[:n])
		if err4 != nil{
			fmt.Println("Error : dsc1.Write = ", err4)
			break
		}else{
			fmt.Println("Info : dsc1.Write 写入成功！")
		}
	}
}
