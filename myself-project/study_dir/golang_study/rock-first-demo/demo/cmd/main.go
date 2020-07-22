package main

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"rockwang.com/rock-first-demo/demo/model"
	"rockwang.com/rock-first-demo/demo/routers"
	"syscall"
)

/*
	1.具有登录注册功能的页面
	2.待定
*/

var rootCmd = &cobra.Command{
	Use:        "demo",
	Short:      "This is a demo website",
	SuggestFor: []string{"demo"}, // 没理解啥用处
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Used to get the version of the application",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("demo version : v1.1.0")
	},
}

var newServerCmd = &cobra.Command{
	Use:   "server",
	Short: "Used to start the demo server",
	Run:   startDemoServer, // start demo server
}

func startDemoServer(cmd *cobra.Command, args []string) {
	// connect database
	DB, err := model.InitDB()
	if err != nil {
		glog.Fatal("Connect database failed, err = ", err)
		return
	}

	// close db connection
	defer model.DBClose(DB)

	router := routers.InitRouter()
	routers.RouterDefine(router)
	// Run routers
	router.Run(":8088")

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	// 顺义 王鑫的用法
	//quit := make(chan os.Signal)
	//signal.Notify(quit, os.Interrupt)
	//<-quit

	// listen signal
	sigs := make(chan os.Signal, 3)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	for {
		select {
		case s := <-sigs:
			glog.Infof("Received signal %v", s)
			os.Exit(0)
		}
	}
}

// Add command to demo server
func init() {
	rootCmd.AddCommand(versionCmd)   // get server version
	rootCmd.AddCommand(newServerCmd) // start demo server
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		glog.Fatal(err)
	}

}
