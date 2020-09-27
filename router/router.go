package router

import (
	"blog/controller"
	"blog/middleware"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func init() {
	f,_:=os.OpenFile("./logs/web.log",os.O_RDWR|os.O_CREATE,0666)
	gin.DefaultWriter = io.MultiWriter(f,os.Stdout)
	router := gin.Default()
	router.Static("/static/avatar","./static/avatar")
	router.Static("/static/image","./static/image")
	router.POST("/blog/register", func(ctx *gin.Context) { controller.Register(ctx) })
	router.POST("/blog/login", func(ctx *gin.Context) { controller.Login(ctx) })
	router.POST("/blog/sendmail", func(ctx *gin.Context) { controller.SendMail(ctx) })
	router.GET("/blog/user/getAll", middleware.TokenParse(), func(ctx *gin.Context) { controller.GetAllUser(ctx) })
	//router.GET("/user/getMine",middleware.TokenParse(),func(ctx *gin.Context){controller.GetMine(ctx)})
	router.Run("localhost:8080")
}
