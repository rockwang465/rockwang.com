package main

import (
	"fmt"
	"net"
)

const (
	defaultMasterCAUrl = "private.ca.sensetime.com"
	defaultSlaveCAUrl  = "slave.private.ca.sensetime.com"
)

// 解析域名为ip
func LookupIP(domain string) ([]string, error) {
	iprecords, err := net.LookupIP(domain)
	ipSlice := make([]string, 0, 10)
	for _, ip := range iprecords {
		//fmt.Println(ip)
		//fmt.Printf("%T\n", ip)
		ipStr := ip.String()
		fmt.Printf("%s, %T\n", ipStr, ipStr)
		ipSlice = append(ipSlice, ipStr)
	}
	return ipSlice, err
}

func main() {
	masterIp, err1 := LookupIP(defaultMasterCAUrl)
	slaveIp, err2 := LookupIP(defaultSlaveCAUrl)
	if err1 != nil {
		fmt.Println("Error: lookup domain err1 :", err1)
	}
	if err2 != nil {
		fmt.Println("Error: lookup domain err2 :", err2)
	}
	//fmt.Printf("master_ip: %#v\n", masterIp)
	//fmt.Printf("slave_ip: %#v\n", slaveIp)
	fmt.Printf("master_ip: %v\n", masterIp)
	fmt.Printf("slave_ip: %v\n", slaveIp)

	licenseMode := ""
	if slaveIp[0] == "127.0.0.1" {
		licenseMode = "standalone"
	} else {
		licenseMode = "distributed"
	}
	fmt.Println("license mode is: ", licenseMode)

}
