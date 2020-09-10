package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"rockwang.com/rock-second-demo/demo/middleWare"
	"rockwang.com/rock-second-demo/demo/model"
	"rockwang.com/rock-second-demo/demo/response"
)

func Login(ctx *gin.Context) {
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

	// 查看数据库是否有此手机号
	DB := model.GetDB()
	var queryUser model.UserInfo
	if DB.Where("telephone = ?", telephone).First(&queryUser).RecordNotFound() {
		response.Response(ctx, http.StatusUnprocessableEntity, 422004, "该手机号未注册", nil)
		return
	}

	// 比对密码
	if err := bcrypt.CompareHashAndPassword([]byte(queryUser.Password), []byte(password)); err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, 422005, "密码不正确", nil)
		return
	}

	// jwt生成token
	tokenStr, err := middleWare.GenerateToken(queryUser.ID)
	if err != nil {
		response.InternalErrResp(ctx, 500002, "", nil)
		glog.Error("Generate token failed , err = ", err)
		return
	}

	// 请求成功，返回信息及token
	response.SuccessResp(ctx, 200001, "登录成功", gin.H{
		"name":          queryUser.Name,
		"telephone":     queryUser.Telephone,
		"Authorization": tokenStr,
	})
}
