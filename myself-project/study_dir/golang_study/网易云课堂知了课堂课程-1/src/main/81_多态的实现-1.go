//package main
//
//import "fmt"
//
//// 1.定义接口
//type OptIF interface {
//	calc() int  // 注意，函数calc()有返回值，这里也要有返回值的定义
//}
//
//type Operator struct {
//	num1 int
//	num2 int
//}
//
//type Sum struct {
//	Operator
//}
//
//type Reduce struct {
//	Operator
//}
//
//// 加法计算
//func (r *Sum) calc() int {
//	return r.num1 + r.num2
//}
//
//// 减法计算
//func (r *Reduce) calc() int {
//	return r.num1 - r.num2
//}
//
//func main() {
//	// 普通用法: 通过对象方法调用
//	//var s1 Sum
//	//s1.num1 = 50
//	//s1.num2 = 20
//	//res := s1.calc()
//	//fmt.Println(res)
//
//	//2.通过接口实现
//	var o OptIF
//	var a Sum = Sum{Operator{10, 20}}
//	o = &a
//	value := o.calc()
//	fmt.Println(value)
//}

package main

import "fmt"

// 1.1 先定义接口
type Humaner81 interface {
	PrintInfo()
}

type student81 struct {
	name  string
	age   int
	score int
}

type teacher81 struct {
	name    string
	age     int
	subject string
}

func (info *student81) PrintInfo() {
	fmt.Printf("大家好，我叫%s， 今年%d岁，我的分数是%d\n", info.name, info.age, info.score)
}

func (info *teacher81) PrintInfo() {
	fmt.Printf("大家好，我叫%s， 今年%d岁，我的学科是%s\n", info.name, info.age, info.subject)
}

// 2.1 多态的实现
// 将接口作为函数参数，来实现多态
func printInformation(h Humaner81) {
	h.PrintInfo()
}

func main() {
	stu := student81{"rock", 18, 98}
	// 3.1 调用多态函数
	printInformation(&stu) // 其实就是调用printInformation()函数而已
}
