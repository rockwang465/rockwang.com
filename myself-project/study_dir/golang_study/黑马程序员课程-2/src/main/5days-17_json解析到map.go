package main

import (
	"encoding/json"
	"fmt"
)

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

	// 定义一个map
	m := make(map[string]interface{}, 4)
	err := json.Unmarshal([]byte(jsonBuf), &m)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	//fmt.Println("m = ", m)     // m =  map[company:itcast isok:true price:66.66 subjects:[go c++ python java]]
	//fmt.Printf("m = %+v\n", m) // m = map[company:itcast isok:true price:66.66 subjects:[go c++ python java]]
	//fmt.Println(m["company"])  // itcast

	// 循环打印，进行断言
	var str string

	for k, v := range m {
		//fmt.Printf("key: %v---->value: %v\n", k, v)
		// 注意，string是无法转换v值的。
		switch  data:= v.(type) {
		case string:
			str = data
			fmt.Printf("map[%s]的值类型是string，value = %s \n", k, str)  // map[company]的值类型是string，value = itcast
		case bool:
			fmt.Printf("map[%s]的值类型是bool，value = %v \n", k, v)  // map[isok]的值类型是bool，value = true
		case float64:
			fmt.Printf("map[%s]的值类型是float64，value = %v \n", k, v)  // map[price]的值类型是float64，value = 66.66
		case []string:
			fmt.Printf("map[%s]的值类型是[]string字符切片，value = %v \n", k, v)  // 这个没有匹配到
		case []interface{}:
			fmt.Printf("map[%s]的值类型是[]interface{}万能空接口切片，value = %v \n", k, v)  // map[subjects]的值类型是[]interface{}万能空接口切片，value = [go c++ python java]
		}

	}
}
