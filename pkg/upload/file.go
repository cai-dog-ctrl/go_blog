package upload

import (
	"blog/global"
	"blog/pkg/util"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

type    FileType  int

const   TypeImage  FileType=  iota  +1  //自增类型的const  ，每一次在新的一行使用时，它的值都会递增，初始化0

// GetFileName 返回格式化后的文件名
func GetFileName (name string)string{
  	ext :=GetFileExt(name)//获取文件类型
  	fileName:=strings.TrimSuffix(name,ext)  //把文件名中后缀名去掉
  	fileName=util.EncodeMD5(fileName)//格式化文件名

  	return fileName+ext
}

// GetFileExt 返回路径使用的文件扩展名,也就是文件类型
func GetFileExt(name string)string{
	return path.Ext(name)
}

// GetSavePath 返回文件路径
func GetSavePath()string{
	return global.AppSetting.UploadSavePath
}

// CheckSavePath 检查保存文件的路径
func CheckSavePath(dst string)bool{
	_,err:=os.Stat(dst)
	return os.IsNotExist(err)
}

// CheckContainExt 检查文件后缀是否包含在约定的后缀配置项中
func CheckContainExt(t FileType,name string)bool{
	ext:=GetFileName(name)
	ext=strings.ToUpper(ext)  //将字母转为大写
	switch t {
	case TypeImage:
		for _,allowExt:=range global.AppSetting.UploadImageAllowExts{
			if strings.ToUpper(allowExt)==ext{
				return true
			}
		}
	}
	return false
}

// CheckMaxSize 检查文件大小
func CheckMaxSize(t FileType,f multipart.File)bool{
	content,_:=ioutil.ReadAll(f)
	size :=len(content)
	switch t {
	case TypeImage:
		if size>=global.AppSetting.UploadImageMaxSize{
			return true
		}
	}
	return false
}

// CheckPermission 检查文件权限是否足够
func CheckPermission(dst string)bool{
	_,err:=os.Stat(dst)
	return os.IsPermission(err)
}

// CreateSavePath 创建在上传文件时所使用的保存目录
func CreateSavePath(dst string,perm os.FileMode)error{
	err:=os.Mkdir(dst,perm)
	if err!=nil{
		return err
	}
	return nil
}

// SaveFile 保存所上传的文件
func SaveFile(file *multipart.FileHeader,dst string)error{
	src,err:=file.Open()
	if err!=nil{
		return err
	}
	defer src.Close()

	out,err:=os.Create(dst)
	if err!=nil{
		return err
	}
	defer out.Close()
	_,err=io.Copy(out,src)
	return err
}