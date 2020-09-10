package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"net/http"
	"rockwang.com/rock-second-demo/demo/commonTools"
	"rockwang.com/rock-second-demo/demo/model"
	"rockwang.com/rock-second-demo/demo/response"
)

func Register(ctx *gin.Context) {
	name := ctx.PostForm("name")
	if len(name) == 0 {
		// random string
		name = commonTools.RandomStr(8)
	}

	password := ctx.PostForm("password")
	if len(password) < 8 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422001, "密码长度必须大于8位", nil)
		return
	}

	telephone := ctx.PostForm("telephone")
	if len(telephone) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422002, "手机号长度必须为11位", nil)
		return
	}

	bytePassword, err := commonTools.Encryption(password)
	if err != nil {
		response.InternalErrResp(ctx, 500003, "", nil)
		glog.Error("encryption password failed , err = ", err)
		return
	}

	DB := model.GetDB()
	user := &model.UserInfo{
		Name:      name,
		Telephone: telephone,
		Password:  string(bytePassword),
	}

	if DB.Where("telephone = ?", telephone).First(&model.UserInfo{}).RecordNotFound() {
		if err := DB.Create(user).Error; err != nil {
			response.InternalErrResp(ctx, 500001, "", nil)
			glog.Error("Add userInfo data failed, err = ", err)
			return
		}
	} else {
		response.Response(ctx, 422, 422003, "该手机号已注册", nil)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"name":      name,
		"telephone": telephone,
		//"password":  bytePassword,
	})
}
