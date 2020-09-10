package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"rockwang.com/rock-first-demo/demo/middlerware"
	"rockwang.com/rock-first-demo/demo/model"
	"strconv"
)

func LoginHandler(ctx *gin.Context) {
	// bind request data to &user
	//var user model.User
	//err := ctx.Bind(&user)
	//if err != nil {
	//	glog.Fatal("Context bind failed : ", err)
	//}

	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")

	// when telephone can not transfer to int type
	_, err := strconv.Atoi(telephone)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": "422001", "msg": "手机号中含有非数字内容",
		})
		return
	}

	// when telephone number length not equal 11
	if len(telephone) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": "422002", "msg": "手机号不是11位",
		})
		return
	}

	// get DB client
	var user model.User
	DB := model.GetDB()
	DB.Where("telephone=?", telephone).First(&user)
	if user.ID == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": "422005", "msg": "该手机号不存在",
		})
		return
	}

	// compare password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": "400001", "msg": "密码错误",
		})
		return
	}

	// create token
	token, err := middlerware.GenerateToken(user.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": "500002", "msg": "系统异常",
		})
		glog.Fatal("token generate failed: ", err)
		return
	}

	// successful login, return token
	ctx.JSON(http.StatusOK, gin.H{
		"code":          "200010",
		"msg":           "登录成功",
		"name":          user.Name,
		"telephone":     telephone,
		"Authorization": token,
	})
}
