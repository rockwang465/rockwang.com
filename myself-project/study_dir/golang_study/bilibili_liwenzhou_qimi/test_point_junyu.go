package main

import "fmt"

type person struct {
	Name string
	Age  int
}

func SetName(p person, name string) {
	p.Name = name
}

func SetNameByPointer(p *person, name string) {
	p.Name = name
}

func (p person) SetPersonName(name string) {
	p.Name = name
}

func (p *person) SetPersonNameByPointer(name string) {
	p.Name = name
}

func main() {
	//p := new(person)
	p := person{"Tom", 12}
	//p.Name = "Tom"
	//p.Age = 12
	fmt.Println(p) // &{Tom 12}

	SetName(p, "Jerry") // 修改失败
	////fmt.Println("address is %p", p)
	fmt.Println("SetName:", p) // &{Tom 12}

	SetNameByPointer(&p, "Jerry")       // 修改成功
	fmt.Println("SetNameByPointer:", p) // &{Jerry 12}

	p.SetPersonName("Tom") // 修改失败
	//p.SetPersonName("Jerry")
	fmt.Println("p.SetPersonName", p) // &{Jerry 12}

	p.SetPersonNameByPointer("Tom")            // 修改成功
	fmt.Println("p.SetPersonNameByPointer", p) // &{Tom 12}
}
