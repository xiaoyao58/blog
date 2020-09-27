package middleware

import (
	"blog/database"
	"blog/entity"
	"blog/utils/logs"
	"github.com/gin-gonic/gin"
	"time"
)

func TokenParse() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		accessToken := ctx.DefaultQuery("access_token","")
		if accessToken == ""{
			accessToken = ctx.PostForm("access_token")
		}
		var tokens []entity.AppToken
		dbErr := database.Db.Select(&tokens, "select * from app_token where access_token=?", accessToken)
		if dbErr != nil {
			ctx.JSON(200, dbErr.Error())
			logs.Error.Println(dbErr.Error())
			panic(dbErr.Error())
		}

		if tokens == nil || tokens[0].ExpireAt<time.Now().Format("2006-01-02 15:04:05"){
			ctx.JSON(200, entity.Throw{Code: 422, Msg: "access_token错误或已过期!"})
			panic("token error")
		}
		var users []entity.User
		db1Err := database.Db.Select(&users, "select * from users where id=?", tokens[0].CreateUser)
		if db1Err != nil {
			panic(db1Err.Error())
		}
		ctx.Set("user",users[0])
		ctx.Next()
	}
}
