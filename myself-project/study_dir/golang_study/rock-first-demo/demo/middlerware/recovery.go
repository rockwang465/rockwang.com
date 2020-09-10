package middlerware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"rockwang.com/rock-first-demo/demo/response"
)

// 返回panic的报错给用户界面，而不是直接被panic了。需要在路由中Use。
func Recovery() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil { // 如果遇到panic，则报错返回给用户
				response.FailResp(ctx, fmt.Sprint(err), nil)
			}
		}()
		ctx.Next()
	}
}
