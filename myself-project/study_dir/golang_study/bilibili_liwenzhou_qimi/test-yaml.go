package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Conf struct {
	Server         Server         `yaml:"server" json:"server"`  // yaml是为了调用yaml方法，必须要加。 json字段为了固定大小写(方便在map中展示)。
	DatabaseSource DatabaseSource `yaml:"databaseSource" json:"database_source"` // 注意，config.yaml中databaseSource和这里的databaseSource一定要相同
}

type Server struct {
	Port        string `yaml:"port" json:"port"`
	ServerIp    string `yaml:"server_ip" json:"server_ip"`
	TokenExpire string `yaml:"token_expire" json:"token_expire"`
}
type DatabaseSource struct {
	Username string `yaml:"username" json:"username"`
	Password string `yaml:"password" json:"password"`
	IP       string `yaml:"ip" json:"ip"`
	Port     string `yaml:"port" json:"port"`
	Database string `yaml:"database" json:"database"`
	Driver   string `yaml:"driver" json:"driver"`
	Charset  string `yaml:"charset" json:"charset"`
	Loc      string `yaml:"loc" json:"loc"`
}

func readInConfig(filePath string) (*Conf, error) {
	config := &Conf{}
	yamlFile, err := ioutil.ReadFile(filePath) // 1.读取文件
	if err != nil {
		return nil, err
	}
	fmt.Println(string(yamlFile))
	err = yaml.Unmarshal(yamlFile, config)  // 2.写入到结构体中
	if err != nil {
		return nil, err
	}

	//fmt.Println("config in readInConfig func:")
	//fmt.Println(config)
	return config, nil
}

func main() {
	config, err := readInConfig("./config.yaml")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	fmt.Println("config:")
	fmt.Println(config)

	confByte, err := json.Marshal(config)  // 3.转为json
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	fmt.Println("confByte:")
	fmt.Println(string(confByte))
	//结果: {"server":{"port":"8090","server_ip":"127.0.0.1","token_expire":"60"},"database_source":{"username":"root","password":"rock1314","ip":"10.151.3.86","port":"3333","database":"demo2","driver":"mysql","charset":"utf8mb4","loc":"AsiaShanghai"}}
}
