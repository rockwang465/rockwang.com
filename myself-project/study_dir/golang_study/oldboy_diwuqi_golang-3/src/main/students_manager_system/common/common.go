package common

import (
	"fmt"
	"strconv"
	"strings"
)

// 检查用户的传参是否正常（1.必须为数字，2.必须为1-5之间的数字）
func (stu *StuManSys) CheckArgs(inputArg string) (isInt bool, intArg int) {
	isInt = false // 是否传入的是数字
	var err error
	intArg, err = strconv.Atoi(strings.Replace(inputArg, "\n", "", -1))
	if err != nil {
		fmt.Println("Error : strconv.Atoi err = ", err)
		return
	}
	isInt = true
	return
}
