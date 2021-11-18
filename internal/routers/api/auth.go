package api

import (
	"blog/global"
	"blog/internal/service"
	"blog/pkg/app"
	"blog/pkg/errcode"
	"github.com/gin-gonic/gin"
)

// GetAuth 校验及获取入参后，绑定并获取到的 app_key 和 app_secret 进行数据库查询，
//检查认证信息是否存在，若存在则进行 Token 的生成并返回。
func GetAuth(c *gin.Context) {
	param := service.AuthRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Infof("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	err := svc.CheckAuth(&param)
	if err != nil {
		global.Logger.Infof("app.CheckAuth errs: %v", errs)
		response.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
		return
	}
	token, err := app.GenerateToken(param.AppKey, param.AppSecret)
	if err != nil {
		global.Logger.Infof("app.GenerateToken errs: %v", errs)
		response.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
		return
	}

	response.ToResponse(gin.H{
		"token": token,
	})
}
