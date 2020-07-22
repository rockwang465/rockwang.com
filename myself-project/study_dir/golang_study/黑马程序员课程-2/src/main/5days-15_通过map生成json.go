package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// 创建一个map, 用interface{}空接口，表示里面可以是任何类型的值(bool,string,int,float都行)
	m := make(map[string]interface{}, 4)
	m["company"] = "itcast"
	m["subjects"] = []string{"go", "python", "c#", "java"}
	m["isok"] = true
	m["price"] = 66.666

	// 编码成json
	res, err := json.Marshal(m)
	//res, err := json.MarshalIndent(m, "", " ")
	if err != nil {
		fmt.Println("err = ,", err)
		return
	}
	fmt.Println(string(res))
	//结果: {"company":"itcast","isok":true,"price":66.666,"subjects":["go","python","c#","java"]}

}
