package main

import (
	"encoding/json"
	"fmt"
)

type IT struct {
	// 3. json二次编码
	//字段名  类型    二次编码
	Company  string   `json:"-"`        //-表示不要输出到屏幕
	Subjects []string `json:"subjects"` // 将大写Subjects改为小写subjects
	IsOk     bool     `json:",string"`  // 转成字符串格式, 例如: "true"
	Price    float64  `json:",string"`  // 转成字符串格式, 例如: "66.66"
}

func main() {
	// 1.定义结构体变量，同时初始化
	s := IT{"itcast", []string{"go", "c++", "python", "java"}, true, 66.66}

	// 2.生成json文本
	// 2.1 普通json格式
	//buf, err := json.Marshal(s)  // string(buf)结果为 {"subjects":["go","c++","python","java"],"string":true,"Price":66.66}

	// 2.2 格式化json格式，输出到屏幕有格式显示，看着更舒服
	buf, err := json.MarshalIndent(s, "", " ")

	if err != nil { // 如果有报错
		fmt.Println("err = ", err)
		return
	}
	fmt.Println(string(buf))
}
