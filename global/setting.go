package global

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-programming-tour-book/blog-service/pkg/logger"
	"github.com/go-programming-tour-book/blog-service/pkg/setting"
)

var (
	ServeSetting    *setting.ServeSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS

	Logger *logger.Logger

	Validator binding.StructValidator
)
