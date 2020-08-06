package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := "1.35 234.2sdf adsf12.34sdf sdf .24 4.52 2dsdf.s dfs.42 24d.23"
	// 获取所有的小数
	res1 := regexp.MustCompile(`\d+\.\d+`)
	if res1 == nil {
		fmt.Println("返回值为空，正则获取失败")
		return
	}else{
		//res2 := res1.FindAllStringSubmatch(str, -1)  // 切片中转为切片
		res2 := res1.FindAllString(str, -1)  // 切片中转为字符串
		fmt.Println(res2)
	}
}
