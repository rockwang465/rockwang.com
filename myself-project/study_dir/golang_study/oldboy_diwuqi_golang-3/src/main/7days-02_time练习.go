package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now) // 2020-06-17 19:13:23.1065918 +0800 CST m=+0.007979501

	// now.Add  加22小时
	n := now.Add(22 * time.Hour)
	fmt.Println(n) // 2020-06-18 17:13:23.1065918 +0800 CST m=+79200.007979501
	//fmt.Println(n.Format("2006-01-02 15:04:05")) // 2020-06-18 17:13:23

	// Sub 两个时间相减
	nextYear, err := time.Parse("2006-01-02 15:04:05", "2019-08-04 12:25:00")
	if err != nil {
		fmt.Println("time Parse err = ", err)
		return
	}
	now = now.UTC()
	//fmt.Println(now) // 2020-06-17 11:13:23.1065918 +0000 UTC
	d := nextYear.Sub(now)
	fmt.Println(d) // -7630h48m23.1065918s
}
