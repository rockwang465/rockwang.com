package middleWare

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 解决跨域问题
func CORSMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		method := ctx.Request.Method
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Max-Age", "86400")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token,Authorization,Token")
		ctx.Header("Access-Control-Allow-Methods", "*") // "POST,GET,OPTIONS"
		ctx.Header("Access-Control-Expose-Headers", "Content-Length,Access-Control-Allow-Origin,Access-Control-Header,Content-Type")
		ctx.Header("Access-Control-Allow-Credentials", "true")
		// 放心所有OPTIONS方法
		if method == "OPTIONS" {
			ctx.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		ctx.Next()
	}
}
