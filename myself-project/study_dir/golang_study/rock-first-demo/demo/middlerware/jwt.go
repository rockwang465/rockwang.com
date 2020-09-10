package middlerware

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/golang/glog"
	"github.com/spf13/viper"
	"strconv"
	"time"
)

// 为用户生成token的中间件

type Claims struct {
	//Username string `json:"username"`
	//Password string `json:"password"`
	UserID uint // 老师这里只传了一个id
	jwt.StandardClaims
}

var jwtKey = []byte("rock_define_a_secret_key") // 随机定义一个字符串即可

// use claims to generate token
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

	expireTime := nowTime.Add(userDefineExpire * time.Minute) // token过期时间定义

	claims := &Claims{
		UserID: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), // 设置token过期时间，这个最重要
			IssuedAt:  nowTime.Unix(),    // token发放时间
			Issuer:    "rock",            // 颁发者
			Subject:   "user token",      //主题
		},
	}

	// token声明
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtKey)
	return token, err
}

// parse token string transfer to claims
func ParseTokenToClaims(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})
	return token, claims, err
}
