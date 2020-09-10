package middlerware

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)

// 解决跨域问题的中间件

func CORSMiddleware() gin.HandlerFunc {
	srvPort := viper.GetString("server.port")
	if srvPort == "" {
		srvPort = "8080"
	}
	localUrl := "http://localhost:" + srvPort
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", localUrl) // localUrl为域名，也可以用*表示所有域名
		ctx.Writer.Header().Set("Access-Control-Max-Age", "86400")       // 设置缓存时间
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "*")     // 设置可以通过访问的方法，可以是: POST GET等， *代表所有
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "*")     // 允许请求携带的信息， *代表所有
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if ctx.Request.Method == http.MethodOptions { // 判断是否为option请求
			ctx.AbortWithStatus(200) // 如果是option请求，则直接返回200 - 为啥abort，防黑客?不知道
		} else {
			ctx.Next() // 否则继续后面的操作
		}
	}
}
