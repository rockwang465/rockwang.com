package main

import "fmt"

type User interface {
	printInfo()
}

type Students struct {
	name  string
	age   int
	sex   string
	class string
	score int
}

func (stu *Students) printInfo() {
	fmt.Printf("Name: %s, age: %d, sex: %s, class: %s , score: %d\n", stu.name, stu.age, stu.sex, stu.class, stu.score)
}

type Teacher struct {
	name    string
	age     int
	sex     string
	project string
}

func (tea *Teacher) printInfo() {
	fmt.Printf("Name: %s, age: %d, sex: %s, project: %s\n", tea.name, tea.age, tea.sex, tea.project)
}

type Manager struct {
	name        string
	age         int
	sex         string
	managerType string
}

func (man *Manager) printInfo() {
	fmt.Printf("Name: %s, age: %d, sex: %s, managerType: %s\n", man.name, man.age, man.sex, man.managerType)
}

func mainInter(u User) {
	u.printInfo()
}

func main() {
	var u User
	stu1 := Students{"rock", 18, "male", "4", 99}
	tea1 := Teacher{"jacky", 28, "male", "English"}
	man1 := Manager{"nova", 23, "female", "teacher and student"}
	u = &stu1
	mainInter(u)
	u = &tea1
	mainInter(u)
	u = &man1
	mainInter(u)
}
