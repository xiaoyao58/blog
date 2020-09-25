package router

import (
	"blog/controller"
	"blog/middleware"
	"github.com/gin-gonic/gin"
)

func init(){
	router := gin.Default()
	router.POST("/blog/register",func(ctx *gin.Context){controller.Register(ctx)})
	router.POST("/blog/sendmail",func(ctx *gin.Context){controller.SendMail(ctx)})
	router.GET("/user/getAll",middleware.TokenParse(),func(ctx *gin.Context){controller.GetAllUser(ctx)})
	//router.GET("/user/getMine",middleware.TokenParse(),func(ctx *gin.Context){controller.GetMine(ctx)})
	router.Run("localhost:8080")
}