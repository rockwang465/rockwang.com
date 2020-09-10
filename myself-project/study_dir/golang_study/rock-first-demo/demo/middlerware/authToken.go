package middlerware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"net/http"
	"rockwang.com/rock-first-demo/demo/model"
	"strings"
)

// 解析用户携带的token是否正确的中间件

func AuthToken(ctx *gin.Context) { // 我的写法，可能不对
	//func AuthToken() gin.HandlerFunc {  // 老师的写法，return func
	//	return func(ctx *gin.Context) {
	// get context token string
	tokenString := ctx.GetHeader("Authorization")

	// check token string
	if !strings.HasPrefix(tokenString, "Bearer ") || len(tokenString) == 0 {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code": "401001", "msg": "权限不足",
		})
		glog.Info("insufficient authority")
		// ctx.Abort() 在被调用的函数中阻止挂起函数。注意这将不会停止当前的函数。
		//例如，你有一个验证当前的请求是否是认证过的 Authorization 中间件。
		//如果验证失败(例如，密码不匹配)，调用 Abort 以确保这个请求的其他函数不会被调用。
		ctx.Abort()
		return
	}

	// parse token string to claims
	tokenString = tokenString[7:] // 去掉前面"Bearer "7个字符
	token, claims, err := ParseTokenToClaims(tokenString)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code": "401002", "msg": "权限不足",
		})
		glog.Warning("parse token string failed : ", err)
		ctx.Abort()
		return
	}

	// invalid token
	if !token.Valid { // 如果不是有效的token，或者有err，就报错
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code": "401003", "msg": "权限不足",
		})
		glog.Warning("token is invalid")
		ctx.Abort()
		return
	}

	// token is valid, get user info
	userId := claims.UserID // 获取用户id
	DB := model.GetDB()
	var user model.User
	DB.Where("id = ?", userId).First(&user)
	if user.ID == 0 {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code": "401004", "msg": "权限不足",
		})
		glog.Info("claims user id not in database :", err)
		return
	}

	// write user info to context
	ctx.Set("user", user)  // 将表中信息保存到context的user key中

	ctx.Next() // 当前中间件结束，执行router命令行中后面的操作
}

//}
