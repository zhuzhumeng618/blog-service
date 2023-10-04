package global

import (
	"blog-service/pkg/logger"
	"blog-service/pkg/setting"
)

var (
	Server   *setting.Server
	App      *setting.App
	Database *setting.Database
	Logger   *logger.Logger
)
