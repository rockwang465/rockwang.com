package controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"rockwang.com/rock-second-demo/demo/commonTools"
	"rockwang.com/rock-second-demo/demo/dto"
	"rockwang.com/rock-second-demo/demo/model"
	"rockwang.com/rock-second-demo/demo/response"
)

func verifyToken(ctx *gin.Context) (*model.UserInfo, error) {
	// 获取token，所有对当前用户的信息进行操作，都是从token中获取userID进行操作(如修改手机号、名称、密码等)，这样更准确
	user, exists := ctx.Get("user") // 获取中间件authToken set保存在context中的user信息
	if !exists {
		return nil, errors.New("Context cat not get [user] key")
	}
	return user.(*model.UserInfo), nil
}

// 展示用户信息
func UserInfo(ctx *gin.Context) {
	user, err := verifyToken(ctx)
	if err != nil {
		response.InternalErrResp(ctx, 403001, "权限不足", nil)
		glog.Info("Context cat not get [user] key")
		return
	}
	// 展示用户信息
	// 先过滤数据，然后展示
	userData := dto.UserDto(user)
	response.SuccessResp(ctx, 200001, "", gin.H{
		"user": userData,
	})
}

// 更新用户昵称
func ChangeName(ctx *gin.Context) {
	user, err := verifyToken(ctx)
	if err != nil {
		response.Response(ctx, http.StatusForbidden, 403001, "权限不足", nil)
		glog.Info("Context cat not get [user] key")
		return
	}

	// 获取用户传入的name
	name := ctx.Params.ByName("name")
	if len(name) == 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422004, "name字段不能为空", nil)
		return
	}

	// 权限检查
	var newUser model.UserInfo
	DB := model.GetDB()
	if DB.Where("id = ?", user.ID).First(&newUser).RecordNotFound() { // 先查询
		response.Response(ctx, http.StatusForbidden, 403001, "权限不足", nil)
		glog.Info("user id record not found 2")
		return
	}

	// 修改用户名
	if err = DB.Model(&newUser).Update("name", name).Error; err != nil { // 再修改
		response.InternalErrResp(ctx, 500002, "更改名称失败，请稍后再试", nil)
		glog.Error("change name failed, err = ", err)
		return
	}

	response.SuccessResp(ctx, 200002, "修改成功", gin.H{
		"new_user": dto.UserDto(&newUser),
	})
}

// 更新用户手机号
func ChangeTel(ctx *gin.Context) {
	user, err := verifyToken(ctx)
	if err != nil {
		response.Response(ctx, http.StatusForbidden, 403001, "权限不足", nil)
		glog.Info("Context cat not get [user] key")
		return
	}

	// 获取用户传入的telephone
	telephone := ctx.Params.ByName("telephone")
	if len(telephone) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422002, "手机号长度必须为11位", nil)
		return
	}

	// 权限检查
	var newUser model.UserInfo
	DB := model.GetDB()
	if DB.Where("id = ?", user.ID).First(&newUser).RecordNotFound() { // 先查询
		response.Response(ctx, http.StatusForbidden, 403001, "权限不足", nil)
		glog.Info("user id record not found 3")
		return
	}

	// 修改手机号
	if err = DB.Model(&newUser).Update("telephone", telephone).Error; err != nil { // 再修改
		response.InternalErrResp(ctx, 500002, "更改名称失败，请稍后再试", nil)
		glog.Error("change name failed, err = ", err)
		return
	}

	response.SuccessResp(ctx, 200002, "修改成功", gin.H{
		"new_user": dto.UserDto(&newUser),
	})
}

// 更新用户密码
func ChangePwd(ctx *gin.Context) {
	// oldpassword newpassword 两个字段进行密码修改

	user, err := verifyToken(ctx)
	if err != nil {
		response.Response(ctx, http.StatusForbidden, 403001, "权限不足", nil)
		glog.Info("Context cat not get [user] key")
		return
	}

	// 获取用户传入的新password
	password := ctx.PostForm("newpassword")
	fmt.Println("new password ->",password)
	if len(password) < 8 {
		fmt.Println(password)
		response.Response(ctx, http.StatusUnprocessableEntity, 422001, "密码长度必须大于8位", nil)
		return
	}

	// 权限检查
	var newUser model.UserInfo
	DB := model.GetDB()
	if DB.Where("id = ?", user.ID).First(&newUser).RecordNotFound() { // 先查询
		response.Response(ctx, http.StatusForbidden, 403001, "权限不足", nil)
		glog.Info("user id record not found 4")
		return
	}

	// 获取用户传入的旧密码,并验证旧密码是否正确
	oldPassword := ctx.PostForm("oldpassword")
	dbPassword := newUser.Password
	if err = bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(oldPassword)); err != nil{
		response.Response(ctx, http.StatusUnprocessableEntity, 422005, "旧密码错误", nil)
		glog.Info("old password is not right")
		return
	}

	// 修改密码
	pwdByte, err := commonTools.Encryption(password)  // 加密
	if err != nil {
		response.InternalErrResp(ctx, 500003, "", nil)
		glog.Error("encryption password failed , err = ", err)
		return
	}
	encrytPwd := string(pwdByte)

	if err = DB.Model(&newUser).Update("password", encrytPwd).Error; err != nil { // 再修改
		response.InternalErrResp(ctx, 500002, "更改名称失败，请稍后再试", nil)
		glog.Error("change name failed, err = ", err)
		return
	}

	response.SuccessResp(ctx, 200002, "修改成功", gin.H{
		"new_user": dto.UserDto(&newUser),
	})
}
