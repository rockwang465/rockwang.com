package routers

import (
	"github.com/gin-gonic/gin"
	"rockwang.com/rock-first-demo/demo/controller"
)

func InitRouter() *gin.Engine {
	// create router
	router := gin.Default()

	return router
	// define router group
}

func RouterDefine(r *gin.Engine) {
	// define router rule
	r.POST("/login", controller.LoginHandler) // 这里LoginHandler代表的是一个函数，所以必须这么些
}
