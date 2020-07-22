package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// 路径可以是绝对路径，也可以是相对路径
	//fp, err := os.Create("a.txt")  // 相对路径
	fp, err := os.Create("./a.txt") // 相对路径
	//fp, err := os.Create("../a.txt")  // 相对路径
	//fp, err := os.Create("D:/a.txt")  //绝对路径
	if err != nil {
		// 1.文件创建失败的原有:
		// A. 路径不存在，无法创建文件夹的
		// B. 文件权限限制
		// C. 程序打开文件上限（fp.Close()忘记用此命令关闭了）
		fmt.Println("文件创建失败")
		return
	} else {
		fmt.Println("文件创建成功")
	}

	// 3.写文件
	// 3.1 常用写文件方法
	//fp.WriteString("hello world !\r\n")  // 常用方式
	//fp.WriteString("hello rock !")
	//fp.WriteString("Thank you !")
	// 注意: go语言中，\n在文件中无法换行。只能用\r\n(回车)换行
	// 因为: windows和linux不同
	// 另外: go写入的汉字是3个字节一个，而windows上手动写入汉字是2个字节一个。

	// 3.2 字符切片写入文件
	//slice := []byte{'h','e','l','l','o'}  // 字符切片用法写入很少用
	//count, err2 := fp.Write(slice)  // 注意，这里是有返回值的,count为切片长度
	//if err2 != nil {
	//	fmt.Println("slice 写入文件失败")
	//}else{
	//	fmt.Println("slice 写入文件成功")
	//	fmt.Println(count)  // 5
	//}

	// 3.3 移动光标写入(和python的seek移动光标一样)
	// os.SEEK_END要被移除了，现在都用io.SeekEnd
	count, _ := fp.Seek(0, io.SeekEnd)
	fmt.Println(count) // 0

	// 指定位置写入
	// 前面为字符切片数据，后面为写入的位置
	fp.WriteAt([]byte("hello world ~~~"), count) // hello world ~~~
	// 如果这里还是从0开始写入，则会覆盖前面的文件
	fp.WriteAt([]byte("gogogo -------"), 0) // gogogo -------~
	fp.WriteAt([]byte("秀儿"), 19)            // 从第19个字节位置开始写入内容

	// 2.关闭文件
	// A.如果打开文件不关闭，造成内存的浪费。
	// B.且影响程序打开文件的上限。
	//fp.Close()
	defer fp.Close() // 建议延迟调用
}
