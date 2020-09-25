package controller

import (
	"blog/entity"
	"blog/service"
	"blog/utils/helper"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

func Register(ctx *gin.Context){
	id := helper.Uuid()
	userName := ctx.PostForm("user_name")
	password := ctx.PostForm("password")
	nickName := ctx.DefaultPostForm("nick_name","")
	realName := ctx.DefaultPostForm("real_name","")
	avatar := ctx.PostForm("avatar")
	registMethod:= ctx.PostForm("regist_method")
	mobile := ""
	email := ""
	if strings.Index(registMethod,"@")>0{
		email = registMethod
	}else{
		mobile = registMethod
	}
	sex := ctx.PostForm("sex")
	birthday := ctx.DefaultPostForm("birthday","")
	createTime := helper.TimeFormat(time.Now().Add(8*time.Hour))
	updateTime := helper.TimeFormat(time.Now().Add(8*time.Hour))
	var user = entity.User {
		Id: id,
		UserName: userName,
		Password: password,
		NickName: nickName,
		RealName: realName,
		Avatar: avatar,
		Mobile: mobile,
		Email: email,
		Sex: sex,
		Birthday: birthday,
		CreateTime: createTime,
		UpdateTime: updateTime,
	}
	service.Register(user,ctx)
	var throw entity.Throw
	throw.Code=201
	throw.Msg="注册成功!"
	ctx.JSON(200,throw)
}

func SendMail(ctx *gin.Context){
	service.SendMail(ctx)
	var throw entity.Throw
	throw.Code = 201
	throw.Msg = "验证码已发送"
	ctx.JSON(200,throw)
}
