package service

import (
	"blog/database"
	"blog/entity"
	"blog/utils"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
)

func Register(user entity.User, ctx *gin.Context) {
	db := database.Db
	var users []entity.User
	var postr entity.PostR
	var numName = 0
	var numMethod = 0
	uchan := make(chan string, 2)

	go func() {
		err := db.Select(&users, "select * from users where user_name=?", user.UserName)
		if err != nil {
			postr.Code = 422
			postr.Msg = err.Error()
			ctx.JSON(200, postr)
			panic(err.Error())
		}
		numName += len(users)
		uchan <- "user_name"
	}()

	go func() {
		err := db.Select(&users, "select * from users where email=? or mobile=?", user.Email, user.Mobile)
		if err != nil {
			postr.Code = 422
			postr.Msg = err.Error()
			ctx.JSON(200, postr)
			panic(err.Error())
		}
		numMethod += len(users)
		uchan <- "regist_method"
	}()

	<-uchan
	<-uchan
	if numName > 0 {
		postr.Code = 403
		postr.Msg = "用户名被占用！"
		ctx.JSON(200, postr)
		panic("用户名被占用！")
	}

	if numMethod > 0 {
		postr.Code = 403
		postr.Msg = "邮箱或手机号已被注册！"
		ctx.JSON(200, postr)
		panic("邮箱或手机号已被注册！")
	}

	r, err_r := redis.Dial("tcp", "127.0.0.1:6379")
	if err_r != nil {
		postr.Code = 422
		postr.Msg = err_r.Error()
		ctx.JSON(200, postr)
		panic(err_r.Error())
	}
	defer r.Close()
	registMethod := ctx.PostForm("regist_method")
	randCode, get_err := redis.String(r.Do("get", registMethod))
	if get_err != nil {
		postr.Code = 422
		postr.Msg = "验证码不正确"
		ctx.JSON(200, postr)
		panic("验证码获取失败")
	}
	verifyCode := ctx.PostForm("verify_code")
	if randCode != verifyCode {
		postr.Code = 422
		postr.Msg = "验证码不正确"
		ctx.JSON(200, postr)
		panic("验证码不正确")
	}


	_, err := db.Exec("insert into users set id=?,user_name=?,password=?,nick_name=?,real_name=?,avatar=?,mobile=?,email=?,sex=?,birthday=?,create_time=?,update_time=?",user.Id,user.UserName,user.Password,user.NickName,user.RealName,user.Avatar,user.Mobile,user.Email,user.Sex,user.Birthday,user.CreateTime,user.UpdateTime )
	if err != nil {
		postr.Code = 422
		postr.Msg = err.Error()
		ctx.JSON(200, postr)
		panic(err.Error())
	}
	return
}


func SendMail(ctx *gin.Context){
	username := ctx.PostForm("user_name")
	email := ctx.PostForm("email")
	fmt.Println(email)
	postr := new(entity.PostR)
	if username=="" || email == ""{
		postr.Code = 422
		postr.Msg = "用户名或邮箱不能为空"
		ctx.JSON(200,postr)
		panic("用户名或邮箱不能为空")
	}
	err := utils.SendMail(email,email)
	if err != nil{
		postr.Code = 500
		postr.Msg = err.Error()
		ctx.JSON(200,postr)
		panic(err.Error())
	}
	return
}
