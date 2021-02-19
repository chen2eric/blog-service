package middleware

import (
	"fmt"
	"time"

	"github.com/chen2eric/blog-service/global"
	"github.com/chen2eric/blog-service/pkg/app"
	"github.com/chen2eric/blog-service/pkg/email"
	"github.com/chen2eric/blog-service/pkg/errcode"
	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defailtMailer := email.NewEmail(&email.SMTPInfo{
			Host:     global.EmailSetting.Host,
			Port:     global.EmailSetting.Port,
			IsSSL:    global.EmailSetting.IsSSL,
			UserName: global.EmailSetting.UserName,
			Password: global.EmailSetting.Password,
			From:     global.EmailSetting.From,
		})
		defer func() {
			if err := recover(); err != nil {
				s := "panic recover err: %v"
				global.Logger.WithCallersFrame().Errorf(s, err)

				err := defailtMailer.SendMail(
					global.EmailSetting.To,
					fmt.Sprintf("异常抛出，发生时间： %d", time.Now().Unix()),
					fmt.Sprintf("错误信息：%v", err),
				)
				if err != nil {
					global.Logger.Panicf("mail.SendMail err: %v", err)
				}
				app.NewResponse(c).ToErrorResponse(errcode.ServerError)
				c.Abort()
			}
		}()
		c.Next()
	}
}
