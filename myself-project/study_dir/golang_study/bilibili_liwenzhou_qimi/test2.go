package main

import (
	"fmt"
	"github.com/golang/glog"
	"io/ioutil"
	"os/exec"
)

func ExecLinuxCommand(strCmd string) (string, error) {
	cmd := exec.Command("/bin/bash", "-c", strCmd)

	stdout, _ := cmd.StdoutPipe()
	if err := cmd.Start(); err != nil {
		//glog.Fatal("Execute failed when Start:" + err.Error())
		return "", err
	}

	outBytes, _ := ioutil.ReadAll(stdout)
	stdout.Close()

	if err := cmd.Wait(); err != nil {
		//fmt.Println("Execute failed when Wait:" + err.Error())
		return "", err
	}
	return string(outBytes), nil
}

func main() {
	//// 获取显卡的卡数
	//nvidiaNumStr, err := ExecLinuxCommand("/usr/bin/nvidia-smi -L | wc -l") // get nvidia card total
	//if err != nil {
	//	glog.Fatal(err)
	//}
	//nvidiaNumStr = strings.Replace(nvidiaNumStr, "\n", "", -1)
	//nvidiaNumInt, err := strconv.Atoi(nvidiaNumStr)
	//fmt.Printf("Execute finished:%d\n", nvidiaNumInt)

	// 安装vps服务

	// 拿到除vps服务卡占用剩余的卡号，且要保证vps占用的数量=pod的数量
	vpsProcess := "video-process-service-worker"
	cmdStrNo := "/usr/bin/nvidia-smi | grep " + vpsProcess + "| awk '{print $2}'"
	cmdStrWc := "/usr/bin/nvidia-smi | grep " + vpsProcess + "| wc -l"
	getVpsNo, err := ExecLinuxCommand(cmdStrNo)     // get nvidia serial number
	getVpsCardWc, err := ExecLinuxCommand(cmdStrWc) // get nvidia card total
	fmt.Println(getVpsNo)
	fmt.Println(getVpsCardWc)

	if err != nil {
		glog.Fatal(err)
	}

	// 等待vps安装完成，查看vps占用的卡情况
	// 如果vps没有安装完成，先拿到stfd tfd ips对应的配置文件
	// 然后卸载stfd tfd ips
	// 然后再次安装vps，等待vps状态
	// 当vps正常了，开始配置stfd tfd ips文件,指定卡号为剩余的卡号
	// 同时安装stfd tfd ips 这3个服务
	// 检测正常/ 结束?
}
