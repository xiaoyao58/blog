package main

import (
	"blog/database"
	_ "blog/router"
)

func main(){
	defer database.Db.Close()
}