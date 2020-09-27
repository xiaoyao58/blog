package service

import (
	"blog/database"
	"blog/entity"
	"blog/utils"
	"blog/utils/logs"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
)

func Register(user entity.User, ctx *gin.Context) {
	db := database.Db
	var users []entity.User
	var throw entity.Throw
	var numName = 0
	var numMethod = 0
	uchan := make(chan string, 2)

	go func() {
		err := db.Select(&users, "select * from users where user_name=?", user.UserName)
		if err != nil {
			throw.Code = 422
			throw.Msg = err.Error()
			ctx.JSON(200, throw)
			logs.Error.Println(err.Error())
			panic(err.Error())
		}
		numName += len(users)
		uchan <- "user_name"
	}()

	go func() {
		err := db.Select(&users, "select * from users where email=? or mobile=?", user.Email, user.Mobile)
		if err != nil {
			throw.Code = 422
			throw.Msg = err.Error()
			ctx.JSON(200, throw)
			logs.Error.Println(err.Error())
			panic(err.Error())
		}
		numMethod += len(users)
		uchan <- "regist_method"
	}()

	<-uchan
	<-uchan
	if numName > 0 {
		throw.Code = 403
		throw.Msg = "用户名被占用！"
		ctx.JSON(200, throw)
		panic("用户名被占用！")
	}

	if numMethod > 0 {
		throw.Code = 403
		throw.Msg = "邮箱或手机号已被注册！"
		ctx.JSON(200, throw)
		logs.Error.Println("邮箱或手机号已被注册！")
		panic("邮箱或手机号已被注册！")
	}

	r, err_r := redis.Dial("tcp", "127.0.0.1:6379")
	if err_r != nil {
		throw.Code = 422
		throw.Msg = err_r.Error()
		ctx.JSON(200, throw)
		logs.Error.Println(err_r.Error())
		panic(err_r.Error())
	}
	defer r.Close()
	registMethod := ctx.PostForm("regist_method")
	randCode, get_err := redis.String(r.Do("get", registMethod))
	if get_err != nil {
		throw.Code = 422
		throw.Msg = "验证码不正确"
		ctx.JSON(200, throw)
		logs.Error.Println("验证码获取失败")
		panic("验证码获取失败")
	}
	verifyCode := ctx.PostForm("verify_code")
	if randCode != verifyCode {
		throw.Code = 422
		throw.Msg = "验证码不正确"
		ctx.JSON(200, throw)
		logs.Error.Println("验证码不正确")
		panic("验证码不正确")
	}


	_, err := db.Exec("insert into users set id=?,user_name=?,password=?,nick_name=?,real_name=?,avatar=?,mobile=?,email=?,sex=?,birthday=?,create_time=?,update_time=?",user.Id,user.UserName,user.Password,user.NickName,user.RealName,user.Avatar,user.Mobile,user.Email,user.Sex,user.Birthday,user.CreateTime,user.UpdateTime )
	r.Do("del",registMethod)
	if err != nil {
		throw.Code = 422
		throw.Msg = err.Error()
		ctx.JSON(200, throw)
		logs.Error.Println(err.Error())
		panic(err.Error())
	}
	return
}


func SendMail(ctx *gin.Context){
	username := ctx.PostForm("user_name")
	email := ctx.PostForm("email")
	throw := new(entity.Throw)
	if username=="" || email == ""{
		throw.Code = 422
		throw.Msg = "用户名或邮箱不能为空"
		ctx.JSON(200,throw)
		logs.Error.Println("用户名或邮箱不能为空")
		panic("用户名或邮箱不能为空")
	}
	err := utils.SendMail(email,email)
	if err != nil{
		throw.Code = 500
		throw.Msg = err.Error()
		ctx.JSON(200,throw)
		logs.Error.Println(err.Error())
		panic(err.Error())
	}
	return
}
