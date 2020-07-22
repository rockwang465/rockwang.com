package main

import (
	"./common"
	"fmt"
	"os"
)

// 此案例来自 https://www.bilibili.com/video/BV19t41157Wn
// 学生管理系统需求
// 1. 打印信息，让用户输入需求
// 2. 登记学生信息
// 3. 修改学生信息
// 4. 打印下所有学生信息
// 5. 退出系统
func main() {
	stuMS := common.NewStuSys()
	stuMS.PrintInfo()
	var inputArg string // 如需表示任意类型，可定义为空接口，
	for {

		// 用户输入选项
		fmt.Printf("\n请输入您需要操作的序号:\n")
		//fmt.Scan(&inputArg) // Rock注意1: 这里inputNo必须加& ，否则不执行

		stdinBuf := make([]byte, 1*1024)
		n, err := os.Stdin.Read(stdinBuf)
		if err != nil {
			fmt.Println("Error : os.Stdin.Read err = ", err)
			continue
		} else {
			inputArg = string(stdinBuf[:n])
		}

		// 判断用户输入选项
		isInt , intArg := stuMS.CheckArgs(inputArg)
		if ! isInt { //如果不是int，则报错
			fmt.Println("Error : 输入的非数字，请重新输入1")
			continue
		}
		switch intArg {
		case 1:
			// 打印选择信息
			stuMS.PrintInfo()
		case 2:
			fmt.Println("登记学员信息")
			stuMS.Register()
		case 3:
			fmt.Println("修改学员信息")
			stuMS.Update()
		case 4:
			stuMS.PrintAllInfo()
		case 5:
			os.Exit(0)
		default:
			fmt.Println("Error : 请输入1-5内选项的数字")
		}
	}
}
