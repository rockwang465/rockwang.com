package common

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type StuManSys struct {
	allStuData []*StuInfo // 结构体切片 [{"rock" 18 "三班" "male"} {"lss" 17 "四班" "female"}]
}

type StuInfo struct {
	studentName string
	age         int
	class       string
	sex         string
}

func NewStuSys() StuManSys {
	return StuManSys{}
}

func (stu *StuManSys) PrintInfo() {
	fmt.Println("===== Welcome use students manager system =====")
	fmt.Println("1. 打印选择信息")
	fmt.Println("2. 登记学员信息")
	fmt.Println("3. 修改学员信息")
	fmt.Println("4. 打印所有学员信息表")
	fmt.Println("5. 退出系统")
	fmt.Println("===============================================")
}

// 1.登记学员信息
func (stu *StuManSys) Register() { // rock注意2: 必须是*StuManSys，非指针数据添加无效，习惯这个用法吧~
	buf := make([]byte, 4*1024)
	fmt.Println("-------- register --------")
	fmt.Println("格式: studentName string, age int, class string, sex string")
	fmt.Println("示例: Rock,18,三班,male")
	fmt.Println("示例: Jacky,19,三班,male")
	fmt.Println("示例: Nova,17,三班,female")
	fmt.Println("-------- register --------")
	fmt.Println("请输入以上格式内容: ")
	n, err := os.Stdin.Read(buf)
	if err != nil {
		fmt.Println("Error : os.Stdin.Read err = ", err)
	}
	regData := strings.Replace(string(buf[:n]), "\n", "", -1)
	regDataList := strings.Split(regData, ",")
	if len(regDataList) != 4 {
		fmt.Println("您登记的学员信息有误，请重新输入")
		fmt.Println("您输入的是: ", regData)
		fmt.Println("请输入此格式: studentName, age, class, sex")
		return
	}
	var tmpStu StuInfo
	tmpStu.studentName = regDataList[0] //本身就是字符串，不需要转(int转string:  strconv.Itoa)
	// rock注意2: err在前面初始化过，这里再用就不用初始化了，否则报错 (err:= 报错)
	tmpStu.age, err = strconv.Atoi(regDataList[1]) // string转int: strconv.Atoi
	if err != nil {
		fmt.Println("Error : strconv.Atoi = ", err)
		return
	}
	tmpStu.class = regDataList[2]
	tmpStu.sex = strings.Replace(regDataList[3], "\n", "", -1)
	stu.allStuData = append(stu.allStuData, &tmpStu) // rock注意: 这里tmpStu要加&

	for i, k := range stu.allStuData {
		fmt.Println(i, *k)
	}
	fmt.Println("登记成功")
}

// 3. 修改学员信息
func (stu *StuManSys) Update() {
	// a.先打印所有学员信息进行展示
	resBool := stu.PrintAllInfo()
	if !resBool { // 如果学员信息数据为空，则不继续执行
		return
	}

	// b.获取用户根据选项选择需要修改的学员信息的id
	fmt.Println("请选择需要修改的学员信息id:")
	buf := make([]byte, 4*1024)
	n, err := os.Stdin.Read(buf)
	if err != nil {
		fmt.Println("Error : os.Stdin.Read err = ", err)
		return
	}
	inputArg := string(buf[:n])
	isInt, intArg := stu.CheckArgs(inputArg)
	if !isInt {
		fmt.Println("Error : 输入的非数字，请重新输入2")
		return
	}
	if intArg >= len(stu.allStuData) || intArg < 0 {
		fmt.Println("Error : 请输入正确范围内的学员id")
		return
	}

	// c.打印用户输入的数字对应的学员信息，若输入正常，则让用户更新学员信息
	fmt.Println(*stu.allStuData[intArg])
	fmt.Println("请输入需要修改信息:")
	fmt.Println("示例: studentName=Lss, age=16, class=五班, sex=female")
	n, err = os.Stdin.Read(buf)
	if err != nil {
		fmt.Println("Error : os.Stdin.Read err = ", err)
		return
	}
	updateArgs := strings.Replace(string(buf[:n]), "\n", "", -1)
	splitArgs := strings.Split(updateArgs, ",")
	if len(splitArgs) > 0 {
		for _, k := range splitArgs {
			//fmt.Println("输入的信息:", k)
			splitCol := strings.Split(k, "=")
			if len(splitCol) == 2 {
				//fmt.Println("当前单个字段为: ", splitCol)
				// 判断当前字段的key是否正常，若正常，则赋值
				switch splitCol[0] {
				case "studentName":
					stu.allStuData[intArg].studentName = splitCol[1]
				case "age":
					v, err := strconv.Atoi(splitCol[1])
					if err != nil {
						fmt.Println("Error : age 不是数字")
					}
					stu.allStuData[intArg].age = v
				case "class":
					stu.allStuData[intArg].class = splitCol[1]
				case "sex":
					stu.allStuData[intArg].sex = splitCol[1]
				default:
					fmt.Println("Error : 不存在此key: ", splitCol[0])
					return
				}
			} else {
				fmt.Println("Error : 错误的值", k)
				return
			}
		}
		fmt.Println("** 操作成功 **")
		fmt.Println(*stu.allStuData[intArg])
	} else {
		fmt.Printf("Error : 您输入的修改信息有误 [%s]\n", updateArgs)
		return
	}

}

// 4. 打印所有学员信息
func (stu *StuManSys) PrintAllInfo() bool {
	if stu.allStuData == nil {
		fmt.Println("您好，当前学员信息表为空")
		return false
	} else {
		for i, subData := range stu.allStuData {
			//fmt.Println(*subData)
			fmt.Printf("No: %d, studentName: %s, age: %d, class: %s, sex: %s\n", i, subData.studentName, subData.age, subData.class, subData.sex)
		}
		return true
	}
}
