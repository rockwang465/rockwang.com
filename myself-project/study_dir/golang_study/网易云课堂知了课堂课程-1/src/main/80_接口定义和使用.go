package main

import "fmt"

// 1.1 定义接口 ，接口命名一般以er结尾，根据接口功能实现
type Humaner interface {
	// 1.2 方法声明，但不是定义方法是怎么实现的
	PrintInfo()
}

type student80 struct {
	name  string
	age   int
	score int
}

type teacher80 struct {
	name    string
	age     int
	subject string
}

func (info *student80) PrintInfo() {
	fmt.Printf("大家好，我叫%s， 今年%d岁，我的分数是%d\n", info.name, info.age, info.score)
}

func (info *teacher80) PrintInfo() {
	fmt.Printf("大家好，我叫%s， 今年%d岁，我的学科是%s\n", info.name, info.age, info.subject)
}

func main() {
	// 2.1 创建一个接口，接口是一种数据类型，可以接收满足对象的信息
	// 接口其实是虚的，方法其实还是调用func中的方法的
	var h Humaner
	stu := student80{"rock", 18, 98}
	//stu.PrintInfo()
	// 2.2 将对象信息赋值给接口类型变量
	h = &stu // 注意: func (info *student80)PrintInfo中*student80是用的地址，所以这里要加& 转为地址
	// 2.3 调用接口
	h.PrintInfo() // 注意，PrintInfo 要和 定义的func PrintInfo 同名

	tea := teacher80{"jack", 38, "math"}
	//tea.PrintInfo()
	h = &tea
	h.PrintInfo()

}
