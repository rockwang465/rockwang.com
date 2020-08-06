package util

import (
	"math/rand"
	"time"
)

// get a random string and return
func GetRandStr(num int) string {
	str := []byte("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	resStr := []byte{}
	rand.Seed(time.Now().UnixNano()) // 时间是一直变化的，所以随机的值也会一直变化
	for i := 0; i < num; i++ {
		//n := rand.Intn(len(str))        // 随机一个数字，随机数为: 0到len(str)-1
		//tmpStr := str[n]                // 拿到对应数字的
		//resStr = append(resStr, tmpStr) // append
		// 简写
		resStr = append(resStr, str[rand.Intn(len(str))])
	}
	return string(resStr)
}
