package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"rockwang.com/rock-first-demo/demo/model"
	"rockwang.com/rock-first-demo/demo/response"
	"rockwang.com/rock-first-demo/demo/util"
	"strconv"
)

func RegisterHandler(ctx *gin.Context) {
	// get DB client
	DB := model.GetDB()

	//ctx := new(gin.Context)

	// get request arguments
	name := ctx.PostForm("name")  //普通获取用户的请求数据 - 可能获取不到，不推荐
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")

	//var requestMap = make(map[string]string)  // 使用map获取用户的请求数据 - 不推荐
	//json.NewDecoder(ctx.Request.Body).Decode(&requestMap)

	//var requestUser = model.User{}  // 通过struct获取用户的请求数据 - 还行
	//json.NewDecoder(ctx.Request.Body).Decode(&requestUser)

	//var requestUser = model.User{}
	//ctx.Bind(requestUser) // 使用Gin框架自带的bind功能获取用户的请求数据 - 推荐(由于我这里没有前端页面，测试发现拿不到请求的数据，只能用前面的方法先用着了)
	//name := requestUser.Name
	//telephone := requestUser.Telephone
	//password := requestUser.Password
	//fmt.Printf("name:[%v,%T],tele:[%v,%T],pwd:[%v,%T]", name, name, telephone, telephone, password, password)

	// when name is empty
	//if name == "" {
	if len(name) == 0 {
		// get a random string, length is 10
		name = util.GetRandStr(10)
	}

	// when telephone number length not equal 11
	if len(telephone) != 11 {
		// StatusUnprocessableEntity, 422: 请求格式正确，但是由于含有语义错误，无法响应
		response.Response(ctx, http.StatusUnprocessableEntity, 422002, "手机号不是11位", nil)
		//ctx.JSON(http.StatusUnprocessableEntity, gin.H{
		//	"code": "422002", "msg": "手机号不是11位",
		//})
		return // exit
	}

	// when telephone can not transfer to int type
	_, err := strconv.Atoi(telephone)
	if err != nil {
		// StatusUnprocessableEntity, 422: 请求格式正确，但是由于含有语义错误，无法响应
		response.Response(ctx, http.StatusUnprocessableEntity, 422001, "手机号中含有非数字内容", nil)
		//ctx.JSON(http.StatusUnprocessableEntity, gin.H{
		//	"code": "422001", "msg": "手机号中含有非数字内容",
		//})
		return // 这里要用return返回用户信息，属于用户问题。 如果用panic、glog.Fatal是属于server端报错，是不该用的。
	}

	// when telephone number has been registry
	// RecordNotFound表示查询结果为空，则返回true
	dbFind := DB.Where("telephone=?", telephone).First(&model.User{}).RecordNotFound()
	if !dbFind { // when not empty
		response.Response(ctx, http.StatusUnprocessableEntity, 422003, "该手机号已注册", nil)
		//ctx.JSON(http.StatusUnprocessableEntity, gin.H{
		//	"code": "422003", "msg": "该手机号已注册",
		//})
		return // exit
	}

	// when password length less than 10
	if len(password) < 10 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422004, "密码不足10位", nil)
		//ctx.JSON(http.StatusUnprocessableEntity, gin.H{
		//	"code": "422004", "msg": "密码不足10位",
		//})
		return // exit
	}

	// encryption password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) // 如果const数字小于默认bcrypt.DefaultCost的const，则把默认的bcrypt.DefaultCost赋值给const，看源码即知。
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500001, "系统异常", nil)
		//ctx.JSON(http.StatusInternalServerError, gin.H{
		//	"code": "500001", "msg": "系统异常",
		//})
		glog.Fatal("bcrypt encryption failed: ", err)
		return
	}

	// save registry info to mysql table
	u := model.User{
		Name:      name,
		Password:  string(hashedPassword),
		Telephone: telephone,
	}
	DB.Create(&u)

	// successful register, return
	response.SuccessResp(ctx, "", gin.H{
		"code":      "200010",
		"msg":       "注册成功",
		"name":      name,
		"telephone": telephone,
	}, )
	//ctx.JSON(http.StatusOK, gin.H{
	//	"code":      "200010",
	//	"msg":       "注册成功",
	//	"name":      name,
	//	"telephone": telephone,
	//})
}
