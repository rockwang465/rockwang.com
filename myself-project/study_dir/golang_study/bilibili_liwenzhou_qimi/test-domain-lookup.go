package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	//dns := "xxbandy.github.io"
	dns := "http://private.ca.sensetime.com"

	// 解析cname
	cname, _ := net.LookupCNAME(dns)

	// 解析ip地址
	ns, err := net.LookupHost(dns)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Err: %s", err.Error())
		return
	}

	// 反向解析(主机必须得能解析到地址)
	//dnsname, _ := net.LookupAddr("127.0.0.1")
	dnsname, _ := net.LookupAddr("10.151.3.94")
	fmt.Println("hostname:", dnsname)

	// 对域名解析进行控制判断
	// 有些域名通常会先使用cname解析到一个别名上，然后再解析到实际的ip地址上
	switch {
	case cname != "":
		fmt.Println("cname:", cname)
		if len(ns) != 0 {
			fmt.Println("vips:")
			for _, n := range ns {
				fmt.Fprintf(os.Stdout, "%s\n", n)
			}
		}
	case len(ns) != 0:
		for _, n := range ns {
			fmt.Fprintf(os.Stdout, "%s\n", n)
		}
	default:
		fmt.Println(cname, ns)
	}
}
