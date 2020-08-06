package main

import "fmt"

// 1.1 先定义接口
type Humaner82 interface {
	PrintInfo()
}

// 2.1 继承接口
type Personer82 interface {
	Humaner82 // 继承了Humaner82接口

	sing(string) //在加一个方法，里面传入字符串
}

type student82 struct {
	name  string
	age   int
	score int
}

type teacher82 struct {
	name    string
	age     int
	subject string
}

func (info *student82) PrintInfo() {
	fmt.Printf("大家好，我叫%s， 今年%d岁，我的分数是%d\n", info.name, info.age, info.score)
}

func (info *teacher82) PrintInfo() {
	fmt.Printf("大家好，我叫%s， 今年%d岁，我的学科是%s\n", info.name, info.age, info.subject)
}

// 3.1 单独创建一个方法，为上面继承接口中的多加的sing唱歌方法
func (s *student82) sing(name string) {
	fmt.Println("我为大家唱首歌，歌名是: ", name)
}

func main() {
	// 4.1 子集及超集的定义
	var h Humaner82  // 子集
	var p Personer82 // 超集(继承子集的接口)

	stu := student82{"rock", 18, 98}
	// 4.2 定义对象
	p = &stu
	p.sing("剑伤")

	// 4.3 接口转换
	//p = h  // error， 注意: 子集是不可以赋值给超集的，
	h = p // 超集可以赋值给子集
	h.PrintInfo()
}
