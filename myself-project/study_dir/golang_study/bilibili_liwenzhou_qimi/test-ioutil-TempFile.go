package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	//file, err := ioutil.TempFile("", "rock-test.txt")
	tmpFile, err := ioutil.TempFile("", "rock-test*")
	if err != nil {
		fmt.Println("Error : ioutil.TempFile failure, ", err)
	}
	defer os.Remove(tmpFile.Name())

	//fmt.Println(file.Name()) // windows: C:\Users\WANGYE~1\AppData\Local\Temp\rock-test187203755
	fmt.Println(tmpFile.Name()) // linux: /tmp/rock-test026205936

	params := []string{}
	params = append(params, fmt.Sprintf("-f %s", tmpFile.Name()))
	p := strings.Join(params, " ")
	fmt.Println(params)
	fmt.Println(p)

	res := fmt.Sprintf("-f %s", tmpFile.Name())
	fmt.Println("rock:")
	fmt.Println(res)
}
