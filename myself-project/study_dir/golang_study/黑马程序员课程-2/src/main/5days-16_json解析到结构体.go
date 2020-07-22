package main

import (
	"encoding/json"
	"fmt"
)

type IT struct {
	//字段名  类型    二次编码
	Company  string   `json:"company"`
	Subjects []string `json:"subjects"`
	IsOk     bool     `json:"isok"`
	Price    float64  `json:"price"`
}

func main() {
	jsonBuf := `
	{
    "company":"itcast",
    "subjects":[
        "go",
        "c++",
        "python",
        "java"
    ],
    "isok":true,
    "price":66.66
	}`

	// 定义一个结构体变量
	var tmp IT
	err := json.Unmarshal([]byte(jsonBuf), &tmp)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Println("tmp = ", tmp) // tmp =  {itcast [go c++ python java] true 66.66}
	fmt.Printf("tmp = %+v\n", tmp)  // tmp = {Company:itcast Subjects:[go c++ python java] IsOk:true Price:66.66}
}
