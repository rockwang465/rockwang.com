package main

import (
	"encoding/json"
	"fmt"
)

//type S struct {
//	name string
//}
//
//func main() {
//	//var m1 map[string]S
//	m1 := make(map[string]S)  // 1.1要初始化，才能进行赋值
//	m1["name1"] = S{"rock"}
//
//	//m2 := map[string]S{"name2": S{"lss"}}  // 2.1 这样是错的，原因是map里结构体无法直接寻址，必须取值
//	m2 := map[string]*S{"name2":&S{"lss"}}  // 2.2 这样才能让下面的操作进行赋值
//	m2["name3"].name = "kevin"  // 2.3 现在赋值不报错了
//}

type Result struct {
	// 这里必须大写，否则不能反序列化
	//status int
	Status int
}

func main() {
	var data = []byte(`{"status":200}`)
	result := &Result{}
	// 进行反序列化
	if err := json.Unmarshal(data, &result); err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("result = %+v\n", result)
	// 小写status结果: result = &{status:0}
	// 大写Status结果: result = &{status:200}
}
