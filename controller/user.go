package controller

import (
	"blog/entity"
	"blog/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetAllUser(ctx *gin.Context){
	fmt.Println(ctx.Query("name"))
	result := service.GetAllUser()
	//json.Marshal(result)
	ctx.JSON(200,result)
}

func GetMine(ctx *gin.Context){
	userIntf,_:=ctx.Get("user")
	user:=userIntf.(entity.User)
	ctx.JSON(200,user)
}
