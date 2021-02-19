package global

import (
	"github.com/chen2eric/blog-service/pkg/logger"
	"github.com/chen2eric/blog-service/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	JWTSetting      *setting.JWTSEttingS
	EmailSetting    *setting.EmailSettingS

	Logger *logger.Logger
)
