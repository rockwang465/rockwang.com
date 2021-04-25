package main

import (
	"github.com/sirupsen/logrus"
	"path"
	"strings"
)

func getConfigNameAndFormat(file string) []string {
	return strings.SplitN(file, ".", 2)
}

func GetConfigName(filePath string) string {
	nameParts := getConfigNameAndFormat(filePath)
	logrus.Infof("nameParts: [%v]", nameParts) // [config yaml]
	return nameParts[0]
}

func main() {
	//configFile := config.GetString("config")
	configFile := "/etc/console/config.yaml"
	fileDir := path.Dir(configFile)        //  文件目录
	logrus.Infof("fileDir: [%v]", fileDir) // /etc/console

	fileFormat := strings.TrimLeft(path.Ext(configFile), ".") // 获取文件名
	logrus.Infof("fileFormat: [%v]", fileFormat)              // yaml

	pathExt := path.Ext(configFile)
	logrus.Infof("pathExt: [%v]", pathExt) // .yaml

	fileBase := path.Base(configFile)
	logrus.Infof("fileBase: [%v]", fileBase) // config.yaml

	fileName := GetConfigName(path.Base(configFile))
	logrus.Infof("fileName: [%v]", fileName) // config
}
