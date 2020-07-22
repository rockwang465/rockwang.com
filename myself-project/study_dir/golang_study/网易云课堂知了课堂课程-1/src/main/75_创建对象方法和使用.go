package main

import "fmt"

type cat struct { //创建一个猫对象
	name string
	age  int
}

type dog struct { //创建一个狗对象
	name string
	age  int
}

// (c cat) 来固定show函数只能被c(猫)对象使用，
// d(狗)对象是用不了的
func (c cat) show() { //为 c 做一个专属的方法
	fmt.Println("喵喵叫")
}

//func (d dog)show(){  //也为 d 做一个同名的专属方法
//	fmt.Println("汪汪汪")
//}

func main() {
	var c cat
	c.name = "小花"
	c.age = 2
	//fmt.Println(c)
	c.show() // 这样，只有c(猫对象)才可以执行show，进行喵喵叫了
	// c为对象，show为方法

	var d dog
	d.name = "旺财"
	d.age = 3
	//fmt.Println(d)
	d.show()
}
