package middleWare

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"net/http"
	"rockwang.com/rock-second-demo/demo/model"
	"rockwang.com/rock-second-demo/demo/response"
)

func AuthToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")

		// parse token type
		if tokenString[0:7] != "Bearer " {
			response.Response(ctx, http.StatusForbidden, 403001, "权限不足", nil)
			glog.Info("Token header is not Bearer ")
			// ctx.Abort() 在被调用的函数中阻止挂起函数。注意这将不会停止当前的函数。
			//例如，你有一个验证当前的请求是否是认证过的 Authorization 中间件。
			//如果验证失败(例如，密码不匹配)，调用 Abort 以确保这个请求的其他函数不会被调用。
			ctx.Abort()
			return
		}

		// parse token
		tokenBody := tokenString[7:]
		token, claim, err := ParseToken(tokenBody)
		if err != nil || ! token.Valid { // error或者token无效
			response.Response(ctx, http.StatusForbidden, 403001, "权限不足", nil)
			if err != nil {
				glog.Info("parse token failed , err = ", err)
			} else {
				glog.Info("Token is invalid ")
			}
			ctx.Abort()
			return
		}

		// 拿到claim中的id进行验证
		userID := claim.UserId
		DB := model.GetDB()
		var user = model.UserInfo{}
		if DB.Where("id = ?", userID).First(&user).RecordNotFound() {
			response.Response(ctx, http.StatusForbidden, 403001, "权限不足", nil)
			glog.Info("user id record not found")
			ctx.Abort()
			return
		}

		// write user info to context
		ctx.Set("user", user) // 将查询的信息放入context的user key中

		ctx.Next()
	}
}
