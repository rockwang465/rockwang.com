package main

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// 4. 定义两个结构体(redis mysql)，分别将ini文件的定义反射到结构体上
type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
	Test     bool   `ini:"test"`
}

type RedisConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Password string `ini:"password"`
	Database int    `int:"database"`
}

// 4.1 将两个结构体放到一个结构体中
type iniConfig struct {
	MysqlConfig `ini:"mysql"`
	RedisConfig `ini:"redis"`
}

func loadIni(fileName string, iCfg interface{}) (err error) {
	// 1.判断传进来的interface类型
	// 1.1 必须是指针类型
	v := reflect.TypeOf(iCfg)
	if v.Kind() != reflect.Ptr { // 判断是否为指针
		fmt.Printf("Error : data not a Ptr type")
		return
	}

	// 1.2 必须是结构体类型
	if v.Elem().Kind() != reflect.Struct { // 判断是否为结构体
		fmt.Printf("Error : data not a Struct type")
		return
	}

	// 2. 读文件
	fileObj, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error : os open mysql.ini failed , err = ", err)
		return
	}
	defer fileObj.Close()

	buf := make([]byte, 1*1024)
	n, err := fileObj.Read(buf)
	if err != nil {
		fmt.Println("Error : read file obj failed , err = ", err)
		return
	}

	// 3. 保存为[]byte格式
	fileRes := string(buf[:n])
	sliceFileRes := strings.Split(fileRes, "\r\n")
	//fmt.Println(fileRes)
	fmt.Printf("%#v\n", sliceFileRes)

	// 5.1 存放存在的结构体(MysqlConfig RedisConfig)
	curStructName := new(string) // 这里必须要初始化，否则只是var的话，后面无法正常引用。

	// 5.2 定义当前的字结(对应的字结mysql redis)
	curSectionName := new(string) // 这里必须要初始化，否则只是var的话，后面无法正常引用。
	saveSection := make(map[string]string)

	// 5 for循环判断文件
	for i, line := range sliceFileRes {

		line = strings.Replace(line, "\r\n", "", -1)
		line = strings.TrimSpace(line)

		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			//fmt.Printf("%d : ; 或 # 开头的注释内容\n", i+1)
			continue
		} else if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") { // 如果是 [ 开头 和 ] 结尾的
			// 去掉头尾，计算长度为0， 则不正常
			resLine := line[1 : len(line)-1]     // 去掉 [ 和 ]
			resLine = strings.TrimSpace(resLine) // 去除首位空格
			if len(resLine) < 1 {
				err = fmt.Errorf("Error : line [%v] syntax error ", i+1)
				fmt.Println(err)
				return
			}
			//fmt.Printf("res is like [xxx] , result : [%s] \n", line)

			// 循环查iniConfig结构体的每一个字段的tag(`ini:"mysql"` `ini:"redis"` 中 tag的key为ini， value为mysql  redis)
			for i := 0; i < v.Elem().NumField(); i++ { // v.NumField() 是反射的一个计算结构体字段数量的方法
				field := v.Elem().Field(i)           // Elem()是转换指针为常量，Field()为获取对应索引的字段
				if field.Tag.Get("ini") == resLine { // field.Tag.Get("ini")  查询当前字段tag的key为ini对应的value值
					fmt.Printf("field.Tag.Get('ini'): [%v] , section: [%v]\n", field.Tag.Get("ini"), resLine)
					*curSectionName = resLine                      // section name : mysql redis
					fmt.Println("field.Name ======> ", field.Name) // MysqlConfig RedisConfig
					*curStructName = field.Name                    // field.Name = MysqlConfig RedisConfig,  string类型
					saveSection[resLine] = field.Name
				}
			}

			// 判断当前的字结section name 是否在map中
			fmt.Println(saveSection)
			if _, ok := saveSection[resLine]; !ok {
				err = fmt.Errorf("Error : not found [%v] section name \n", resLine)
				fmt.Println(err)
				return
			}
		} else {
			// 判断: a.字符串为空 continue
			if line == "" {
				//fmt.Printf("Info : 空 [%v], index: %v \n", line, i+1)
				continue
				// 判断: b.=不在头尾
			} else if strings.HasPrefix(line, "=") || strings.HasSuffix(line, "=") {
				err = fmt.Errorf("Error : line [%v] syntax error ", i+1)
				fmt.Println(err)
				return

			} else {
				//fmt.Printf("Info : [%v], index: %v \n", line, i+1)
				// 判断: c.以=分割，len为2
				sliceLine := strings.Split(line, "=")

				if len(sliceLine) != 2 {
					err = fmt.Errorf("Error : line [%v] syntax error ", i+1)
					fmt.Println(err)
					return
				} else { // 最后正确的字段都到这里了
					fmt.Printf("Info : 最后正确的字段都到这里了 [%v]\n", line)

					// 转为valuesOf进行判断
					v := reflect.ValueOf(iCfg)                        // Rock注意1: 这里要取值，所以要用ValueOf才行
					structObj := v.Elem().FieldByName(*curStructName) // 传入结构体中的字段名(MysqlConfig RedisConfig)
					structTypeInfo := structObj.Type()

					// 判断 *curStructName(MysqlConfig 和 RedisConfig) 为一个结构体
					if structObj.Kind() != reflect.Struct {
						fmt.Printf("Error : %s not struct type\n", *curStructName)
					}

					// 拿到line的 key和 value
					key := strings.TrimSpace(sliceLine[0])
					value := strings.TrimSpace(sliceLine[1])

					var fieldName string
					var fieldType reflect.StructField
					// for循环拿到结构体中每一个字段
					for i := 0; i < structObj.NumField(); i++ {
						field := structTypeInfo.Field(i) // 拿到类型操作对象
						// 如果Tag == key，则保存相应信息
						if field.Tag.Get("ini") == key { // 找到key和结构体字段相同的field字段
							fieldName = field.Name
							fieldType = field
							break // 找到key对应的字段后跳出循环
						}
					}
					// 这里， key == tag的，所以准备赋值给结构体中的key
					if len(fieldName) == 0 {
						// 在结构体中找不到对应的字符
						continue
					}
					fieldObj := structObj.FieldByName(fieldName) // 进行fieldName字段操作

					// 开始赋值，将value赋值给结构体中的key
					fmt.Println(fieldName, fieldType.Type.Kind()) // 这里能判断mysql.ini每行的value的类型
					switch fieldType.Type.Kind() {
					case reflect.String:
						fieldObj.SetString(value) // 如果是String类型，则赋值
					case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
						var valueInt int64 // 为了下面err不用 := 赋值，这样return就能正常使用err返回值了
						// 将字符串转换为数字的函数,功能灰常之强大,看的我口水直流.
						// strconv.ParseInt(value, 10 , 64)中， 10为十进制， 64为int64
						valueInt, err = strconv.ParseInt(value, 10, 64) // strconv.ParseInt()返回int64 和error
						if err != nil {
							err = fmt.Errorf("line: %d value type not int", i+1)
							return
						}
						fieldObj.SetInt(valueInt)
					case reflect.Bool:
						var valueBool bool
						valueBool, err = strconv.ParseBool(value)
						if err != nil {
							err = fmt.Errorf("line: %d value type not bool", i+1)
						}
						fieldObj.SetBool(valueBool)
					case reflect.Float32, reflect.Float64:
						var valueFloat float64
						valueFloat, err = strconv.ParseFloat(value, 64)
						if err != nil {
							err = fmt.Errorf("line: %d value type not float", i+1)
							return
						}
						fieldObj.SetFloat(valueFloat)
					}
				}
			}
		}
	}
	return
}

func main() {
	var iCfg iniConfig                   // Rock注意2: 这里var，否则后面的判断全都错
	err := loadIni("./mysql.ini", &iCfg) // Rock注意2: 这里传入指针，否则后面的判断全都错
	if err != nil {
		fmt.Println("loadIni func err = ", err)
		return
	}
	fmt.Println(iCfg) // {{10.20.30.40 3306 root abc123} {127.0.0.1 6379 abc 123 0}}
	//fmt.Printf("%#v\n", iCfg) // main.iniConfig{MysqlConfig:main.MysqlConfig{Address:"10.20.30.40", Port:3306, Username:"root", Password:"abc123"}, RedisConfig:main.RedisConfig{Host:"127.0.0.1", Port:6379, Password:"abc 123", Database:0}}
}
