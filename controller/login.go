package controller

import (
	"blog/service"
	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context){
	result:=service.Login(ctx)
	ctx.JSON(200,result)
}
