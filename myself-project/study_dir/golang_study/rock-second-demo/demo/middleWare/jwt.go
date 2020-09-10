package middleWare

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/golang/glog"
	"github.com/spf13/viper"
	"strconv"
	"time"
)

type Claim struct {
	UserId uint
	jwt.StandardClaims
}

var jwtKey = []byte("Is_a_jwt_secret_key_from_rock")

// generate a token
func GenerateToken(id uint) (string, error) {
	nowTime := time.Now()
	strDefineExpire := viper.GetString("server.token_expire")
	var intDefineExpire int
	var err error
	if len(strDefineExpire) == 0 {
		glog.Warning("Warning: not defined the server token expire time in application.yaml ")
		fmt.Println("Warning: not defined the server token expire time in application.yaml ")
		intDefineExpire = 30
	}else {
		intDefineExpire, err = strconv.Atoi(strDefineExpire)
		if err != nil {
			glog.Warning("transfer to int failed, err = ", err)
			return "", nil
		}
	}

	userDefineExpire := time.Duration(intDefineExpire)  // 时间格式转换
	//var intDefineExpire time.Duration = 30
	expireTime := nowTime.Add(time.Minute * userDefineExpire)

	claim := &Claim{
		UserId: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), // token过期时间。Unix()将时间转为int64类型的秒数
			IssuedAt:  nowTime.Unix(),    // token发放时间
			Issuer:    "Rock Wang",
			Subject:   "Login token",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenStr, err := tokenClaims.SignedString(jwtKey)
	return tokenStr, err
}

// parse token body
func ParseToken(tokenString string) (*jwt.Token, *Claim, error) {
	var claims = &Claim{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})
	return token, claims, err
}
