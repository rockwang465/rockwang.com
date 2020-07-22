package main

import "fmt"

type person76 struct {
	id   int
	name string
}

func (s person76) Print() { // 给person76一个名字，名字叫s。这里的s相当于函数中的实参，你改为任何字符串都是一样的含义，不要理解错了
	fmt.Printf("%s", s.name)
}

func main() {
	stu := person76{1, "rock"}
	// 解释: 这里的stu就是代表了Print前面的s。 所以，stu可以使用person76的Print方法
	stu.Print()

	//s := person76{2, "jack"}
	//s.Print()
}
