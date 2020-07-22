package main

import "fmt"

type person78 struct {
	id   int
	name string
	age  int
}

type student78 struct {
	person78
	class string
}

// 父类的专属方法: PrintInfo
func (p *person78) PrintInfo() {
	fmt.Printf("id: %d\n", p.id)
	fmt.Printf("姓名: %s\n", p.name)
	fmt.Printf("年龄: %d\n", p.age)
}

// 子类的专属方法: 也叫 PrintInfo
func (s *student78) PrintInfo() {
	fmt.Println("student Info:", *s)
}

func main() {
	s := student78{person78{1, "rock", 18}, "9班"}
	s.PrintInfo() //  打印内容: student Info: {{1 rock 18} 9班}
	// s调用父类和子类同名的方法时，使用就近原则，会先用自己的方法
}
