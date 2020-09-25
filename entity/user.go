package entity

type User struct{
	Id string `db:"id",json:"id"`
	UserName string `db:"user_name"`
	Password string `db:"password"`
	NickName string `db:"nick_name"`
	RealName string `db:"real_name"`
	Avatar string `db:"avatar",json:"avatar"`
	Mobile string `db:"mobile"`
	Email string `db:"email"`
	Sex string `db:"sex"`
	Birthday string `db:"birthday"`
	CreateTime string `db:"create_time"`
	UpdateTime string `db:"update_time"`
}


