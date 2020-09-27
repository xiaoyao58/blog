package service

import (
	"blog/database"
	"blog/entity"
	"blog/utils/logs"
	"github.com/gin-gonic/gin"
)

func GetAllUser() []gin.H{
	var users []entity.User
	err:= database.Db.Select(&users,"select * from users")
	if err!=nil{
		logs.Error.Println("get users error:"+err.Error())
	}
	var user_list []gin.H
	for _,u:=range users{
		f:=gin.H{
			"id":u.Id,
			"user_name":u.UserName,
			"password": u.Password,
			"nick_name":u.NickName,
			"real_name":u.RealName,
			"avatar":u.Avatar,
			"mobile":u.Mobile,
			"email":u.Email,
			"sex": u.Sex,
			"birthday":u.Birthday,
		}
		user_list=append(user_list, f)
	}
	return user_list
}
