package routers

import (
	"blog/global"
	"blog/internal/service"
	"blog/pkg/app"
	"blog/pkg/convert"
	"blog/pkg/errcode"
	"blog/pkg/upload"
	"github.com/gin-gonic/gin"
)

type Upload struct {}

func NewUpload()Upload{
	return Upload{}
}

func (u Upload)UploadFile(c *gin.Context){
	response:=app.NewResponse(c)//新建一个响应体
	file ,fileHeader,err:=c.Request.FormFile("file")
	if err!=nil{
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}
	fileType:=convert.StrTo(c.PostForm("type")).MustInt()
	if fileHeader==nil||fileType<=0{
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	svc:=service.New(c.Request.Context())
	fileInfo,err:=svc.UploadFile(upload.FileType(fileType),file,fileHeader)
	if err!=nil{
		global.Logger.Infof("svc.UploadFile err: %v", err)
		response.ToErrorResponse(errcode.ErrorUploadFileFail.WithDetails(err.Error()))
		return
	}
	response.ToResponse(gin.H{
		"file_access_url":fileInfo.AccessUrl,
	})
}