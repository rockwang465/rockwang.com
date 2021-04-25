package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	addr := "127.0.0.1:8082"

	userGroup := router.Group("/user")
	{
		userGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"user": "/user/index",
			})
		})
		userGroup.GET("/list", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"user": "/user/list",
			})
		})
		userGroup.GET("/info", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"user": "/user/info",
			})
		})
	}

	shopGroup := router.Group("/shop")
	{
		shopGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"shop": "/shop/index",
			})
		})
		shopGroup.GET("/list", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"shop": "/shop/list",
			})
		})
		shopGroup.GET("/info", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"shop": "/shop/info",
			})
		})
	}

	router.Run(addr)
}
