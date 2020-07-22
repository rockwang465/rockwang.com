package main

import (
	"encoding/json"
)

type Server struct {
	ServerName string
	ServerIP   string
}
type Serverslice struct {
	Servers []Server
}

func main() {
	var s Serverslice

	// 模拟传输的Json数据
	str := `{
	"servers": [
                    {
                        "serverName": "Shanghai_VPN",
                        "serverIP": "127.0.0.1"
                    }, {
                        "serverName": "Beijing_VPN",
                        "serverIP": "127.0.0.2"
                    }
                ]
            }`

	// 解析字符串为Json
	json.Unmarshal([]byte(str), &s)

	// 遍历Json
	for key, val := range s.Servers {

		// 打印索引和其他数据
		print(`Key：`, key, "\t")
		print(`Name：`, val.ServerName, "\t")
		println(`IP：`, val.ServerIP)
	}
}
