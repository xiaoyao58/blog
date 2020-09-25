package database

import (
	"blog/config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)
var Db *sqlx.DB
func init(){
	mysql := config.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", mysql.UserName, mysql.Password, mysql.Host, mysql.Port, mysql.DbName, mysql.Charset)
	database, err := sqlx.Open("mysql", dsn)
	if err != nil{
		log.Fatal(err.Error())
	}
	Db = database
	Db.SetMaxIdleConns(30)
	Db.SetMaxOpenConns(30)
}
