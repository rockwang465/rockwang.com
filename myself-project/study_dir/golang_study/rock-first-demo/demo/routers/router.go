package routers

import (
	"github.com/gin-gonic/gin"
	"rockwang.com/rock-first-demo/demo/controller"
	"rockwang.com/rock-first-demo/demo/middlerware"
)

func InitRouter() *gin.Engine {
	// create router
	router := gin.Default()

	return router
}

func RouterDefine(r *gin.Engine) {
	// use middleware
	r.Use(middlerware.CORSMiddleware(), middlerware.Recovery())

	// define router rule
	r.POST("/register", controller.RegisterHandler) // 1.这里RegisterHandler代表的是一个函数，所以必须这么些
	r.POST("/login", controller.LoginHandler)       // 2.登录函数

	r.GET("/user/info", middlerware.AuthToken, controller.UserInfo) // 3.用户信息页面，增加token认证中间件

	// define article category router group
	category := r.Group("/article_categories")
	articleCategoryRoutes := controller.NewCategory()
	{
		category.POST("", articleCategoryRoutes.Create) // postman中使用id请求
		category.GET("/:id", articleCategoryRoutes.Show)
		category.PUT("/:id", articleCategoryRoutes.Update)
		category.DELETE("/:id", articleCategoryRoutes.Delete)

	}

	// defind article rules 单篇文章操作
	post := r.Group("/post")
	post.Use(middlerware.AuthToken)
	postController := controller.NewPostController()
	{
		post.POST("", postController.Create)
		post.GET(":uuid", postController.Show)
		post.PUT(":uuid", postController.Update)
		post.DELETE(":uuid", postController.Delete)
		// 注意: 如果是GET请求的page list，则会和GET请求的show方法中":uuid"冲突的
		post.POST("page/list", postController.PageList) // 分页
	}

}
