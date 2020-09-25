package entity

import "time"

type AppToken struct{
	Id string `db:"id"`
	CreateUser string `db:"create_user"`
	AccessToken string `db:"access_token"`
	ExpireAt time.Time `db:"expire_at"`
	CreateAt time.Time `db:"creat_at"`
	UpdateAt time.Time `db:"modify_at"`
	ClientInfo string `db:"client_info"`
}