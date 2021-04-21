package main

import (
	//"flag"
	"github.com/golang/glog"
)

func main() {
	//flag.Parse()

	//defer glog.Flush()

	glog.Info("This is info message")
	glog.Infof("This is info message: %v", 12345)
	glog.InfoDepth(1, "This is info message", 12345)

	glog.Warning("This is warning message")
	glog.Warningf("This is warning message: %v", 12345)
	glog.WarningDepth(1, "This is warning message", 12345)

	glog.Error("This is error message")
	glog.Errorf("This is error message: %v", 12345)
	glog.ErrorDepth(1, "This is error message", 12345)

	//glog.Fatal("This is fatal message")
	//glog.Fatalf("This is fatal message: %v", 12345)
	//glog.FatalDepth(1, "This is fatal message", 12345)

	glog.V(1).Infoln("level 1")     // 3
	glog.V(2).Infoln("level 2")
	//defer glog.Flush()
	glog.Flush()
}