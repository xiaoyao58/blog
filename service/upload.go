package service

import (
	"blog/utils/helper"
	"github.com/gin-gonic/gin"
	"strings"
)

//单文件上传:
//[参数1:*gin.Context,参数2:保存目录]
//[返回1:文件名称,返回2:错误信息]
func Upload(ctx *gin.Context, destDir string) (string, error) {
	file, err1 := ctx.FormFile("file")
	if err1 != nil {
		return "", err1
	}
	FileName := file.Filename
	fileName := helper.Uuid() + FileName[strings.LastIndex(FileName, "."):]
	dest := destDir + "/" + fileName
	err2 := ctx.SaveUploadedFile(file, dest)
	if err2 != nil {
		return "", err2
	}
	return fileName, nil
}


//多文件上传:
//[参数1:*gin.Context,参数2:保存目录]
//[返回1:文件名称列表,返回2:错误信息]
func Multipart(ctx *gin.Context,destDir string )([]string,error){
	fileList := make([]string,0,10)
	mult,err1 := ctx.MultipartForm()
	if err1 !=nil{
		return fileList,err1
	}
	files := mult.File["files"]
	for _,f:=range files{
		FileName := f.Filename
		fileName := helper.Uuid() + FileName[strings.LastIndex(FileName, "."):]
		dest := destDir + "/" + fileName
		err2 := ctx.SaveUploadedFile(f,dest)
		if err2 !=nil {
			return fileList,err2
		}
		fileList = append(fileList,fileName)
	}
	return fileList,nil
}
