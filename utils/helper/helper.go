package helper

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

func TimeFormat(t time.Time) string{
	res := t.Format("2006-01-02 15:04:05")
	return res
}

func Uuid() (myuuid string){
	myuuid = uuid.NewV4().String()
	return
}







