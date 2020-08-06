package main

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// 这里纯Rock自己手写的
// 1.生成url地址
func httpGet(url string, i int, GXStatus chan<- int) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("http.Get err = ", err) // 这里就不报错了，把错误返回到外面就行
		return
	}

	buf := make([]byte, 4*1024)
	var res string

	for {
		n, _ := resp.Body.Read(buf)
		res += string(buf[:n]) // 由于返回值中有res，相当于已经自动做了一个 var res string的声明了
		if n == 0 {            // 读完退出for循环
			break
		}
	}

	defer resp.Body.Close()
	//fmt.Println(res)

	// 2.过滤有用内容
	filterRespRes := filterRespBody(res)

	// 3.将爬到的内容写入到文件中
	fileName := "gaoxiaoba" + strconv.Itoa(i) + ".html"
	e1, e2 := writeToFiles(fileName, filterRespRes)
	if e1 != nil {
		fmt.Println("writeTofiles err1 = ", e1)
	}
	if e2 != nil {
		fmt.Println("writeTofiles err2 = ", e2)
	}

	// 难点2: 这里是最后操作部分，直接传入内容即可
	GXStatus <- i // channel是全局的，所以这里传进去后，其他地方就能读到了
}

// 2.过滤有用内容
func filterRespBody(respBody string) (res string) {
	rule := `最近流行</a>.+`
	compile := regexp.MustCompile(rule)
	var subMatch []string
	if compile == nil {
		fmt.Println("regexp.MustCompile err")
	} else {
		subMatch = compile.FindAllString(respBody, -1)
		//fmt.Println("subMatch :",subMatch)
	}

	// 讲<div>开头替换成换行符\n
	res = strings.Replace(subMatch[0], "<div", "\n<div", -1)
	return

}

// 3.将爬到的内容写入到文件中
func writeToFiles(fileName string, filterRespRes string ) (err1 error, err2 error) {
	fp, err1 := os.Create(fileName)
	if err1 != nil {
		//fmt.Println("os.Create err = ", err1)
		return
	}

	//fmt.Println(respBody)  // 测试这里是有结果的
	_, err2 = fp.WriteString(filterRespRes)
	if err2 != nil {
		return
	}

	return // 最后是这里return，别弄错了
}

func main() {
	// 贴吧高校吧
	baseUrl := "http://tieba.baidu.com/f/index/forumpark?pcn=%E9%AB%98%E7%AD%89%E9%99%A2%E6%A0%A1&pci=0&ct=1&rn=20&pn="

	GXStatus := make(chan int)

	// 老师这里的i的值，是通过fmt.Scan()获取用户输入从1-25页的，这样两个都是变量了。开始start(1)，结束end(25)，更好用
	for i := 1; i < 25; i++ {
		// 1.生成url地址
		gaoXiaoUrl := baseUrl + strconv.Itoa(i) // i 要转成字符串后才能拼接
		go httpGet(gaoXiaoUrl, i, GXStatus)
	}

	for i := 1; i < 25; i++ {
		fmt.Printf("已经完成第[%d]次页面爬取\n", <-GXStatus)  // 难点1: 要这里打印结果，消耗掉管道内的内容，让下一次go协程继续被卡住
	}

	fmt.Println("http get over")
}
