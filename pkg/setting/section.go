package setting

import (
	"time"
)

//  声明配置属性的结构体

type SeverSettingS struct {
	RunMode string
	HttpPort string
	ReadTimeout time.Duration
	WriteTimeout time.Duration
}

type AppSettingS struct {
	DefaultPageSize int
	MaxPageSize int
	LogSavePath string
	LogFileName string
	LogFileExt string
	UploadSavePath       string
	UploadServerUrl      string
	UploadImageMaxSize   int
	UploadImageAllowExts []string
}

type DatabaseSettingS struct {
	DBType       string
	UserName     string
	Password     string
	Host         string
	DBName       string
	TablePrefix  string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

// ReadSection 获取一个键并将其解组为一个结构体
func (s *Setting)ReadSection (k string,v interface{})error{
	err:=s.vp.UnmarshalKey(k,v)
	if err!=nil{
		return err
	}
	return nil
}
