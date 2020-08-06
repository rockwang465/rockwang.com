package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"net/http"
	"rockwang.com/rock-first-demo/demo/model"
	"rockwang.com/rock-first-demo/demo/util"
	"strconv"
)

func LoginHandler(ctx *gin.Context) {
	// get DB client
	DB, err := model.InitDB()
	if err != nil {
		glog.Fatal(err)
	}
	// get name
	//ctx := new(gin.Context)
	name := ctx.PostForm("name")
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")
	// when name is empty
	//if name == "" {
	if len(name) == 0 {
		// get a random string, length is 10
		name = util.GetRandStr(10)
	}

	// when telephone can not transfer to int type
	_, err = strconv.Atoi(telephone)
	if err != nil {
		// StatusUnprocessableEntity, 422: 请求格式正确，但是由于含有语义错误，无法响应
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": "810", "msg": "手机号中含有非数字内容", "current_length": len(telephone),
		})
		return // 这里要用return返回用户信息，属于用户问题。 如果用panic、glog.Fatal是属于server端报错，是不该用的。
	}

	// when telephone number length not equal 11
	if len(telephone) != 11 {
		// StatusUnprocessableEntity, 422: 请求格式正确，但是由于含有语义错误，无法响应
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": "811", "msg": "手机号不是11位", "current_length": len(telephone),
		})
		return // exit
	}

	// when telephone number has been registry
	// RecordNotFound表示查询结果为空，则返回true
	dbFind := DB.Where("telephone=?", telephone).First(&model.Demo{}).RecordNotFound()
	if !dbFind { // when not empty
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": "812", "msg": "该手机号已注册",
		})
		return // exit
	}

	// when password length less than 10
	if len(password) < 10 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": "813", "msg": "密码不足10位", "current_length": len(password),
		})
		return // exit
	}

	// save registry info to mysql table
	u := model.Demo{
		Name:      name,
		Password:  password,
		Telephone: telephone,
	}
	DB.Create(&u)
	//fmt.Println("resCre 创建结果:", resCre)

	// successful login
	ctx.JSON(http.StatusOK, gin.H{
		"code": "800", "msg": "注册成功", "name": name, "telephone": telephone, "password": password,
	})
}
