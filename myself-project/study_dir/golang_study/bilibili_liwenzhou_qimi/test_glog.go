package main

import "flag"
import "github.com/golang/glog"

func main() {
	//初始化命令行参数
	flag.Parse()

	//退出时调用，确保日志写入文件中
	defer glog.Flush()

	glog.Info("hello, glog")
	glog.Warning("warning glog")
	glog.Error("error glog")

	glog.Infof("info %d", 1)
	glog.Warningf("warning %d", 2)
	glog.Errorf("error %d", 3)
}
