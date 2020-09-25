package config

type MySql struct {
	UserName string
	Password string
	Host     string
	Port     int
	DbName   string
	Charset  string
}

var Mysql = MySql{
	UserName: "用户名",
	Password: "密码",
	Host:     "主机",
	Port:     3306,
	DbName:   "数据库",
	Charset:  "utf8",
}
