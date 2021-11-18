package global

import (
	"blog/pkg/logger"
	"blog/pkg/setting"
)

var (
	ServerSetting   = new(setting.SeverSettingS)
	AppSetting      = new(setting.AppSettingS)
	DatabaseSetting = new(setting.DatabaseSettingS)
	Logger          = new(logger.Logger)
	JWTSetting      = new(setting.JWTSetting)
)
