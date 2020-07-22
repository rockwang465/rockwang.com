package main

import "fmt"

// 1.定义接口
type OptIF interface {
	calc() int // 注意，函数calc()有返回值，这里也要有返回值的定义
}

type Operator struct {
	num1 int
	num2 int
}

type Sum struct {
	Operator
}

type Reduce struct {
	Operator
}

// 加法计算
func (r *Sum) calc() int {
	return r.num1 + r.num2
}

// 减法计算
func (r *Reduce) calc() int {
	return r.num1 - r.num2
}

// 2.多态的实现
func calc(o OptIF) { //这里函数名和上面calc故意同名，当然不同名也可以
	res := o.calc()
	fmt.Println(res)
}

func main() {
	//// a.普通用法: 通过对象方法调用
	//var s1 Sum
	//s1.num1 = 50
	//s1.num2 = 20
	//res := s1.calc()
	//fmt.Println(res)

	//// b.通过接口实现
	//var o OptIF
	//var a Sum = Sum{Operator{10, 20}}
	//o = &a
	//value := o.calc()
	//fmt.Println(value)

	// 3.通过多态的实现
	var a Sum = Sum{Operator{10, 20}}
	calc(&a) // 其实就是调用calc(o OptIF)这个函数
}
