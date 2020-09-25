package entity

type AppToken struct{
	Id string `db:"id"`
	CreateUser string `db:"create_user"`
	AccessToken string `db:"access_token"`
	ExpireAt string `db:"expire_at"`
	CreateAt string `db:"creat_at"`
	ModifyAt string `db:"modify_at"`
	ClientInfo string `db:"client_info"`
}