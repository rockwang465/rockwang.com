package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Response(ctx *gin.Context, httpCode, returnCode int, notice string, data gin.H) {
	if notice == "" {
		ctx.JSON(httpCode, gin.H{
			"code": returnCode,
			"data": data,
		})
	} else if data == nil {
		ctx.JSON(httpCode, gin.H{
			"code": returnCode,
			"msg":  notice,
		})
	} else {
		ctx.JSON(httpCode, gin.H{
			"code": returnCode,
			"msg":  notice,
			"data": data,
		})
	}
}

// 200 OK
func SuccessResp(ctx *gin.Context, returnCode int, notice string, data gin.H) {
	if len(notice) == 0 {
		notice = "请求成功"
	}
	Response(ctx, http.StatusOK, returnCode, notice, data)
}

// 500 内部错误 StatusInternalServerError
func InternalErrResp(ctx *gin.Context, returnCode int, notice string, data gin.H) {
	if notice == "" {
		notice = "服务端异常，请稍后再试"
	}
	Response(ctx, http.StatusInternalServerError, returnCode, notice, data)
}
