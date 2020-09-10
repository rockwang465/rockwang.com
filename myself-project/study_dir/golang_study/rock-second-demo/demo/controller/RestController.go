package controller

import "github.com/gin-gonic/gin"

type RestController interface {
	Create(ctx *gin.Context) // 此为ArticleCategory结构体的4中方法
	Update(ctx *gin.Context)
	Show(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

