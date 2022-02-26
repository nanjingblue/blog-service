package global

import (
	"github.com/go-programing-tour-book/blog-service/pkg/logger"
	"github.com/go-programing-tour-book/blog-service/pkg/setting"
	"github.com/jinzhu/gorm"
)

var (
	ServerSetting   *setting.ServerSettings
	AppSetting      *setting.AppSettings
	DatabaseSetting *setting.DatabaseSettings
	JWTSetting      *setting.JWTSettings
	EmailSetting    *setting.EmailSettings
	DBEngine        *gorm.DB
	Logger          *logger.Logger
)
