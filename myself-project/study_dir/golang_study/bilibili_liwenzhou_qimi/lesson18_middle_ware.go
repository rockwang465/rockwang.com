package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// 正常的路由中写就是HandlerFunc
func indexHandler(c *gin.Context) {
	fmt.Println("Index func")
	name, ok := c.Get("name") // 5.接收设置的传值
	if !ok {
		name = "匿名用户"
	}

	// 也可以用c.MustGet
	c.JSON(http.StatusOK, gin.H{
		"username": name,
	})
}

// 定义第一个中间件
func m1(c *gin.Context) {
	fmt.Println("m1 in...") // m1 进入
	start := time.Now()

	c.Next() // 调用后续的处理函数
	//c.Abort() // 阻止调用后续的处理函数

	// 计算耗时
	cost := time.Since(start)
	//end := time.Now()
	//cost := end.Sub(start)
	fmt.Printf("cost: %v\n", cost)
	fmt.Println("m1 out ...") // m1 结束
}

// 定义第二个中间件
func m2(c *gin.Context) {
	fmt.Println("m2 in...") // m2 进入
	c.Set("name", "q1mi")   // 4.设置传值: name=q1mi

	//c.Abort() // 阻止调用后续的处理函数
	fmt.Println("m2 out ...") // m2 结束
}

// 认证中间件
func authMiddleWare(doCheck bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		if doCheck { // true，今天过节，不需要登录认证，大家都可以访问电影页面了
		} else {
			// 存放具体的逻辑
			// 判断用户是否登录
			c.Next() // 正常大家都需要登录认证，或者是VIP才能看
		}
	}
}

func main() {
	addr := "127.0.0.1:8080"
	router := gin.Default()

	// 1.全局注册中间件函数m1
	router.Use(m1, m2, authMiddleWare(true))

	// 2.或者将中间件放路由的请求中(这里可以传入多个中间件)
	//router.GET("/index", m1, m2, indexHandler)
	router.GET("/index", indexHandler)
	router.GET("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"user": "ok",
		})
	})
	router.GET("/shop", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"shop": "ok",
		})
	})

	// 3.路由组中增加中间件
	videoGroup := router.Group("/video", authMiddleWare(true))
	{
		videoGroup.GET("/movie", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"movie": "ok",
			})
		})
		videoGroup.GET("/cartoon", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"cartoon": "ok",
			})
		})
	}

	router.Run(addr)
}

// 结果:
//m1 in...
//m2 in...
//m2 out ...
//Index func
//cost: 0s
//m1 out ...
