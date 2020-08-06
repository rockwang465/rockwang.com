package main

import "fmt"

func main() {
	info := "上海自来水来自海上"
	//info := "abcdefgfedcba"
	// 由于info字符串为中文，正常截取索引的值为是乱码，所以要给他换成字符进行比对。
	// 正常字符串换成字符用[]byte(info),但由于这里info是中文，所以得用[]rune(info),这是专为中文定义的字符切片方法。
	//china_word := make([]rune, len(info))
	lenInfo := len([]rune(info))
	b := true
	for i, v := range []rune(info) {
		//head := v
		tail := string([]rune(info)[lenInfo-1-i])
		if string(v) != tail {
			fmt.Println("不是回文")
			b = false
			break
		}
	}
	if b == true {
		fmt.Println("是回文")
	}
}
