package logs

import (
	"log"
	"os"
)

var Error *log.Logger
var Info *log.Logger


func init() {
	f, err := os.OpenFile("./logs/errors.log", os.O_TRUNC|os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err !=nil{
		log.Panic(err.Error())
	}
	e := log.New(f, "[error]", log.LstdFlags|log.Llongfile)
	Error = e


	f1, err1 := os.OpenFile("./logs/infos.log", os.O_TRUNC|os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err1 !=nil{
		log.Panic(err1.Error())
	}
	i := log.New(f1, "[info]", log.LstdFlags|log.Llongfile)
	Info = i
}
