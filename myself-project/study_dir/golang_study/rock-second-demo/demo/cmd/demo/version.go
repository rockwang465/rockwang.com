package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:     "version",
	Short:   "Show demo version",
	Example: "demo version",
	Version: "v1.0.0",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("demo version: v1.0.0")
	},
}
// 也可用写成函数方式
//func VersionCmd() *cobra.Command {
//	versionCmd := &cobra.Command{
//		Use:     "version",
//		Short:   "Show demo version",
//		Example: "demo version",
//		Version: "v1.0.0",
//		Run: func(cmd *cobra.Command, args []string) {
//			fmt.Println("demo version: v1.0.0")
//		},
//	}
//	return versionCmd
//}
