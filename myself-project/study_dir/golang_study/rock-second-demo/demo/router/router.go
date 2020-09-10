package router

import (
	"github.com/gin-gonic/gin"
	"rockwang.com/rock-second-demo/demo/controller"
	"rockwang.com/rock-second-demo/demo/middleWare"
)

type Routers struct {
	Router *gin.Engine
}

func NewRouter() *Routers {
	r := gin.Default()
	r.Use(middleWare.CORSMiddleware(), middleWare.Recovery())
	return &Routers{Router: r}
}

// 注册及登录
func (r *Routers) RegisterAndLogin() {
	r.Router.POST("/register", controller.Register)
	r.Router.POST("/login", controller.Login)
}

// 用户信息页面
func (r *Routers) UserInfo() {
	user := r.Router.Group("/user")
	user.GET("/info", middleWare.AuthToken(), controller.UserInfo)
	user.PUT("/info/change_name/:name", middleWare.AuthToken(), controller.ChangeName)
	user.PUT("/info/change_tel/:telephone", middleWare.AuthToken(), controller.ChangeTel)
	user.PUT("/info/change_pwd", middleWare.AuthToken(), controller.ChangePwd)
}

// 文章分类的增删改查操作
func (r *Routers) ArticleCategory() {
	categoryController := controller.NewArticleCategory() // 实例化文章分类的操作函数
	ac := r.Router.Group("/article_categories")
	{
		ac.GET("/info/:id", middleWare.AuthToken(), categoryController.Show)
		ac.POST("/info", middleWare.AuthToken(), categoryController.Create)
		ac.PUT("/info/:id", middleWare.AuthToken(), categoryController.Update)
		ac.DELETE("/info/:id", middleWare.AuthToken(), categoryController.Delete)
		ac.GET("/all_info", middleWare.AuthToken(), categoryController.AllInfo)
	}
}

// 单篇文章的增删改查操作(post-文章)
func (r *Routers) Post() {
	postController := controller.NewPostController() // 实例化文章分类的操作函数
	r.Router.Use(middleWare.AuthToken()) // 这个一定要放在r.Router.Group的上面
	pc := r.Router.Group("/posts")
	{
		pc.GET("/:uuid", postController.Show)
		pc.POST("", postController.Create)
		pc.PUT("/:uuid", postController.Update)
		pc.DELETE("/:uuid", postController.Delete)
		pc.POST("/page/list", postController.PageList)
	}
}
