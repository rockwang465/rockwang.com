package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rockwang.com/rock-first-demo/demo/dto"
	"rockwang.com/rock-first-demo/demo/model"
)

func UserInfo(ctx *gin.Context) {
	value, boolRes := ctx.Get("user") // 取authToken中保存到上下文的user信息
	if !boolRes {
		ctx.JSON(http.StatusOK, gin.H{
			"code": "200", "msg": "can not get user info",
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": "200", "msg": dto.UserInfoResponse(value.(model.User)), // value.(model.User)强制转换数据类型。value为保存的user信息
	})
	// value的信息太多，特别是password ，所以这里要过滤一些，不要把敏感信息都返回了。
	// {
	//    "code": "200",
	//    "msg": {
	//        "ID": 2,
	//        "CreatedAt": "2020-07-25T16:07:23+08:00",
	//        "UpdatedAt": "2020-07-25T16:07:23+08:00",
	//        "DeletedAt": null,
	//        "Name": "rock1",
	//        "Password": "$2a$10$4775fkqQV0s/TrGWYcWWuODt1ge5VZxPOKSpqG2AqYXngkXDGqIXi",
	//        "Telephone": "17521004201"
	//    }
	//}
}
