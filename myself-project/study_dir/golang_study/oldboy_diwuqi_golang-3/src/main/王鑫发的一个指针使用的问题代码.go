package main

import "fmt"

type student struct {
	name string
	age  int
}

func main() {
	m := make(map[string]*student) // m = ["aa": {name: "小王子", age: 18} ]
	stus := []student{
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "小可爱", age: 9000},
	}
	//fmt.Println(stus) // Rock认为: stus = [{name: "小王子", age: 18},	{name: "娜扎", age: 23},{name: "小可爱", age: 9000}]
	//                  实际打印结果: stus = [{小王子 18} {娜扎 23} {小可爱 9000}]

	for _, stu := range stus { // stu = {小王子 18}
		//m[stu.name] = &stu  //rock注意1: 原代码这里是这样写的，如果这样写会导致打印的信息全是 "小可爱"
		//                      然后将赋值问题换成下面，中间加了一个tmpStu，再赋值给m[stu.name]
		//            解释原因:
		//                因为&stu指向的地址永远都是相同的，所以第二次会覆盖第一次的值，第三次会覆盖第二次的值。
		//            解决方法:
		//                1.value的值不要用指针。但毕竟value是结构体，不用指针还是不太好。
		//                2.用下面的方法，用个新的变量(下面的tmpStu)接收value，然后再用指针赋值过去。
		tmpStu := stu
		m[stu.name] = &tmpStu // m[stu.name]即m[小王子] = &{小王子 18}
		fmt.Println(stu.name, &stu)
	}

	//fmt.Println(m) // Rock认为: m = map[小王子: &{小王子 18} , 娜扎: &{娜扎 23} , 小可爱: &{小可爱 9000}]
	//               实际打印结果: m = map[小王子:0xc0000044a0 娜扎:0xc0000044a0 小可爱:0xc0000044a0 ]
	//               rock注意2:  这里可以看到，如果用上面的原代码，这里的指针地址完全都是一样的。

	for k, v := range m { // k =
		//fmt.Println(k, "k<---->*v", *v)

		fmt.Println(k, "=>", v.name)
		//结果却是:
		//小王子 => 小可爱
		//娜扎 => 小可爱
		//小可爱 => 小可爱
	}
}
