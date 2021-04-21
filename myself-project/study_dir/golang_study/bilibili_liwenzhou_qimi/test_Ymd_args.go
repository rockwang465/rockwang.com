package main

import (
	"fmt"
	"path/filepath"
)

func main(){
	logFile := filepath.Join("/opt/", "rock.%Y%m%d.log")
	fmt.Println(logFile) // \opt\rock.%Y%m%d.log
	// 所以这里是不能识别出来的，只有程序能识别出来
}