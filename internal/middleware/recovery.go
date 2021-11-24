package middleware

import (
	"blog/global"
	"blog/pkg/app"
	"blog/pkg/email"
	"blog/pkg/errcode"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

// Recovery recovery中间件
func Recovery() gin.HandlerFunc {
	//实体化一个Email，并且写入信息
	defailtMailer := email.NewEmail(&email.SMTPInfo{
		Host:     global.EmailSetting.Host,
		Port:     global.EmailSetting.Port,
		IsSSL:    global.EmailSetting.IsSSL,
		UserName: global.EmailSetting.UserName,
		Password: global.EmailSetting.Password,
		From:     global.EmailSetting.From,
	})
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				//当recover捕获了panic，
				global.Logger.WithCallersFrames().Infof("panic recover err: %v", err)

				err := defailtMailer.SendMail(
					global.EmailSetting.To,
					fmt.Sprintf("异常抛出，发生时间: %d", time.Now().Unix()),
					fmt.Sprintf("错误信息: %v", err),
				)
				if err != nil {
					//输出错误日志
					global.Logger.Infof("mail.SendMail err: %v", err)
				}
				//将错误返回给响应体
				app.NewResponse(c).ToErrorResponse(errcode.ServerError)
				c.Abort()
			}
		}()
		//没错的话，就继续执行中间件之后的
		c.Next()
	}
}
