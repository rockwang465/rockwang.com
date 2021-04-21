package main

import (
	"fmt"
	"os/exec"
)

const (
	PythonVenv = "/root/venv/bin/python3"
	SetNetworkCommand = "/usr/local/bin/set-network"
)

var (
	OldIpAddr = "10.151.3.108"
	//OldIpAddr = "10.151.3.99"
	NewIpAddr = "10.151.3.99"
	//NewIpAddr = "10.151.3.108"
	Interface = "eth0"
	Gateway   = "10.151.3.254"
	Netmask   = "255.255.255.0"
)

func main() {
	err := ExecSetNetwork()
	if err != nil {
		panic(err)
	}
}

func ExecSetNetwork() error {
	// /root/venv/bin/python3 /usr/local/bin/set-network --iface=eth0 --old-ip=10.151.3.108 --new-ip=10.151.3.99 --gateway=10.151.3.254 --netmask=255.255.255.0
	cmdArgs := []string{
		SetNetworkCommand,
		"--old-ip", OldIpAddr,
		"--new-ip", NewIpAddr,
		"--iface", Interface,
		//"--gateway", Gateway,
		//"--netmask", Netmask,
	}

	if Gateway != "" {
		cmdArgs = append(cmdArgs, "--gateway", Gateway)
	}
	if Netmask != "" {
		cmdArgs = append(cmdArgs, "--netmask", Netmask)
	}
	fmt.Println("CmdArgs:")
	fmt.Printf("%#v\n", cmdArgs)

	fmt.Println("old_ip:", OldIpAddr)
	fmt.Println("new_ip:", NewIpAddr)
	fmt.Println("interface:", Interface)
	fmt.Println("gateway:", Gateway)
	fmt.Println("netmask:", Netmask)

	//if err := exec.Command(SetNetworkCommand, cmdArgs...).Start(); err != nil {
	if err := exec.Command(PythonVenv, cmdArgs...).Start(); err != nil {
		return fmt.Errorf("start set network failed: %v", err)
	}
	return nil
}
