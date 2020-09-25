package utils

import (
	"blog/database"
	"blog/entity"
	"blog/utils/helper"
	uuid "github.com/satori/go.uuid"
	"time"
)

func GenerateAccessToken(userId string,mytoken chan string) {
	c:=make(chan string)
	var tokens [] entity.AppToken
	db := database.Db
	go func (){
		err:= db.Select(&tokens,"select * from app_token where create_user=? and expire_at>=?",userId,time.Now())
		if err != nil{
			throw:=new(entity.Throw)
			throw.Code=422
			throw.Msg = err.Error()
			panic(err.Error())
			return
		}
		if len(tokens) <=0{
			c<-""
		}else{
			c<-tokens[0].AccessToken
		}
	}()

	randCode:=RandCode("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789",12)
	new_uuid := uuid.NewV5(uuid.NewV4(),userId).String()
	new_uuid = new_uuid[24:26]+new_uuid[26:27]+new_uuid[27:28]+new_uuid[28:29]+new_uuid[32:33]+new_uuid[36:]+randCode[1:2]+randCode[2:3]+randCode[5:6]+randCode[7:8]+randCode[8:9]+randCode[12:]

	isToken:= <-c
	if isToken != ""{
		mytoken<-isToken
		return
	}else{
		go func (){
			_,err:=db.Exec("delete from app_token where create_user=?",userId)
			if err !=nil{
				panic(err.Error())
				return
			}
			token := entity.AppToken{
				Id: helper.Uuid(),
				CreateUser: userId,
				AccessToken: new_uuid,
				ExpireAt: helper.TimeFormat(time.Now().Add(32*time.Hour)),
				CreateAt: helper.TimeFormat(time.Now().Add(8*time.Hour)),
				UpdateAt: helper.TimeFormat(time.Now().Add(8*time.Hour)),
				ClientInfo: "",
			}
			_,erri:=db.Exec("insert into app_token set id=?,create_user=?,access_token=?,expire_at=?,client_info=?,create_at=?,update_at=?",
				token.Id,token.CreateUser,token.AccessToken,token.ExpireAt,token.ClientInfo,token.CreateAt,token.UpdateAt)
			if err != nil{
				panic(erri.Error())
			}
		}()
		mytoken<-new_uuid
		return
	}
}
