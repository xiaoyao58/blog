package service

import (
	"blog/database"
	"blog/entity"
	"blog/utils"
	"github.com/gin-gonic/gin"
)

// @Title 用户登录
// @Param user_name formData string true "用户名或邮箱或手机号"
// @Param password formData string true "密码"
// @Success 200 用户登录成功
// @Failure 500 用户登录失败
// @router /login [post]
func Login(ctx *gin.Context) gin.H {
	username := ctx.PostForm("user_name")
	password := ctx.PostForm("password")
	var users []entity.User

	db := database.Db
	err := db.Select(&users, "select * from users where (user_name=? or email=? or mobile=?) and password=?", username, username, username, password)
	num := len(users)
	if err != nil {
		throw := entity.Throw{
			Code: 422,
			Msg:  err.Error(),
		}
		ctx.JSON(200, throw)
		panic(err.Error())
	}
	if num <= 0 {
		throw := entity.Throw{
			Code: 404,
			Msg:  "用户名或密码错误",
		}
		ctx.JSON(200, throw)
		panic("用户名或密码错误")
	}

	u := users[0]
	mytoken := make(chan string)
	go utils.GenerateAccessToken(u.Id, mytoken)
	token := <-mytoken

	g := gin.H{
		"access_token": token,
		"user_name":    u.UserName,
		"avatar":       u.Avatar,
		"emali":        u.Email,
		"mobile":       u.Mobile,
	}
	return g
}
