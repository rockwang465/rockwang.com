package commonTools

import (
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"strconv"
	"time"
)

func RandomStr(num int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	strBytes := []byte(str)
	resStr := []byte{}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < num; i++ {
		v := rand.Intn(len(str))
		resStr = append(resStr, strBytes[v])
	}
	return string(resStr)
}

// 加密password
func Encryption(password string) ([]byte, error) {
	bytePassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return []byte{}, err
	}
	return bytePassword, nil
}

// 字符串转int
func StrTransferInt(str string) (int, error) {
	intVal, err := strconv.Atoi(str)
	return intVal, err
}
