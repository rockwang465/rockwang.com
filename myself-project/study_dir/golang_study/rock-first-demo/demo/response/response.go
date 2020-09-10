package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Encapsulating response data 封装响应数据
func Response(ctx *gin.Context, httpCode, code int, msg string, data gin.H) {
	if msg == "" {
		ctx.JSON(httpCode, gin.H{
			"code": code,
			"data": data,
		})
	} else if data == nil {
		ctx.JSON(httpCode, gin.H{
			"msg":  msg,
			"code": code,
		})
	} else {
		ctx.JSON(httpCode, gin.H{
			"code": code,
			"msg":  msg,
			"data": data,
		})
	}
}

// Success response data
func SuccessResp(ctx *gin.Context, msg string, data gin.H) {
	Response(ctx, http.StatusOK, 200, msg, data)
}

// Bad response data
func FailResp(ctx *gin.Context, msg string, data gin.H) {
	Response(ctx, http.StatusBadRequest, 400, msg, data)
}
