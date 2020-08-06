package main

import "fmt"

type person79 struct {
	id   int
	name string
	age  int
}

func (p person79) PrintInfo1() {
	fmt.Printf("%p, %v\n", &p, p)
}

func (p *person79) PrintInfo2() {
	fmt.Printf("%p, %v\n", &p, *p)
}

func main() {
	p1 := person79{1, "rock", 18}
	//p1.PrintInfo1()  // 0xc0000044c0, {1 rock 18}
	//p1.PrintInfo2()  // 0xc000006030, {1 rock 18}

	//p1.PrintInfo1  // 为此函数的地址
	//fmt.Printf("%T\n", p1.PrintInfo1)  // func() 类型，确认是函数地址了
	//fmt.Printf("%T\n", p1.PrintInfo2)  // func() 类型，确认是函数地址了

	// 1. 方法值: 隐式传递，隐藏的是接受者，绑定实例(对象)
	// var p1func1 func()  // 第一种定义方法值的办法
	//p1func1 := p1.PrintInfo1  //第二种定义方法值的办法
	//p1func1()

	// 2. 方法表达式
	p2func1 := person79.PrintInfo1
	p2func1(p1)
	p2func2 := (*person79).PrintInfo2
	p2func2(&p1)
}
