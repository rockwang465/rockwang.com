package main

import "fmt"

type person77 struct {
	id   int
	name string
	age  int
}

type student77 struct {
	person77        // 1.使用person78的机构体
	class    string // 2.学生多一个班级元素
}

// 3. 为p制作专属方法
func (p *person77) PrintInfo() {
	fmt.Printf("id: %d\n", p.id)
	fmt.Printf("姓名: %s\n", p.name)
	fmt.Printf("年龄: %d\n", p.age)
}

func main() {
	//p := person77{1,"rock", 18}
	//p.PrintInfo()
	s := student77{person77{2, "jack", 20}, "2班"}
	s.PrintInfo() // 这里s继承了person77的PrintInfo方法
}
