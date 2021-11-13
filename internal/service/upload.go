package service

import (
	"blog/global"
	"blog/pkg/upload"
	"errors"
	"mime/multipart"
	"os"
)

type FileInfo struct {
	name string
	AccessUrl string
}

// UploadFile 上传文件
func (svc *Service)UploadFile(fileType upload.FileType,file multipart.File,fileHeader *multipart.FileHeader)(*FileInfo,error){
	fileName:=upload.GetFileName(fileHeader.Filename)
	if !upload.CheckContainExt(fileType,fileName){
		return nil,errors.New("file suffix is not supported")
	}
	if upload.CheckMaxSize(fileType,file){
		return nil,errors.New("exceeded maximum file limit")
	}
	uploadSavePath:=upload.GetSavePath()
	if upload.CheckSavePath(uploadSavePath){
		if err:=upload.CreateSavePath(uploadSavePath,os.ModePerm);err != nil {
			return nil,errors.New("failed to create save directory")
		}
	}
	if upload.CheckPermission(uploadSavePath){
		return nil, errors.New("insufficient file permissions")
	}
	//前面全是判断能不能上传
	dst := uploadSavePath + "/" + fileName
	if err:=upload.SaveFile(fileHeader,dst);err!=nil{
		return nil,err
	}

	accessUrl:=global.AppSetting.UploadServerUrl+"/"+fileName
	return &FileInfo{name: fileName,AccessUrl: accessUrl},nil
}
