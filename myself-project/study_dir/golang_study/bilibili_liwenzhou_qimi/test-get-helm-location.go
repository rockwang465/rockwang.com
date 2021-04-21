package main

import (
	"fmt"
	"path/filepath"
	"runtime"
)

func main() {
	binLocation := GetBinFileLocation()
	fmt.Println(binLocation)
}

func GetBinFileLocation() string {
	BinName := ""
	switch runtime.GOOS {
	case "windows":
		BinName = "helm.exe"
		return filepath.Join("console", "tools", BinName)
	default:
		BinName = "helm"
	}
	return BinName
}
