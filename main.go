package main

import (
	"blog-service/global"
	"blog-service/internal/model"
	"blog-service/internal/routers"
	"blog-service/pkg/logger"
	"blog-service/pkg/setting"
	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"time"
)

func init() {
	// 程序启动时初始化配置
	if err := setupSetting(); err != nil {
		log.Fatalf("init.setipSetting err: %v", err.Error())
	}

	// 程序启动时初始化数据源
	if err := setupDBEngine(); err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err.Error())
	}

	// 程序启动时初始化日志
	if err := setupLogger(); err != nil {
		log.Fatalf("init.setupLogger err: %v", err.Error())
	}
}

// @title 博客系统
// @version 1.0
// @description go 编程之旅
// @termsOfService http://localhost:8080
func main() {
	// 配置 gin 运行环境
	gin.SetMode(global.Server.RunMode)
	// 获取 gin 实例
	router := routers.NewRouter()
	// 自定义 HTTP 配置
	server := &http.Server{
		Addr:           ":" + global.Server.HttpPort, // gin 服务端口
		Handler:        router,                       // gin ServeHTTP
		ReadTimeout:    global.Server.ReadTimeout,    // gin 读取请求超时时间
		WriteTimeout:   global.Server.WriteTimeout,   // gin 写入响应超时时间
		MaxHeaderBytes: 1 << 20,                      // gin 请求头的最大字节数
	}

	// 使用自定义配置启动 HTTP 监听服务
	if err := server.ListenAndServe(); err != nil {
		panic(err.Error())
	}
}

// setupSetting 初始化 setting 实例
func setupSetting() error {
	// 获取配置实例
	set, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = set.ReadSection("Server", &global.Server)
	if err != nil {
		return err
	}
	err = set.ReadSection("App", &global.App)
	if err != nil {
		return err
	}
	err = set.ReadSection("Database", &global.Database)
	if err != nil {
		return err
	}

	// 配置文件中时间处理
	global.Server.ReadTimeout *= time.Second
	global.Server.WriteTimeout *= time.Second

	return nil
}

// setupDBEngine 初始化 DBEngine 实例
func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.Database)
	if err != nil {
		return err
	}
	return nil
}

// setupLogger 初始化 Logger 实例
func setupLogger() error {
	global.Logger = logger.NewLogger(
		&lumberjack.Logger{
			Filename: global.App.LogSavePath + "/" + global.App.LogFileName + global.App.LogFileExt,
			MaxSize:  600, // 最大大小，MB
			MaxAge:   10,  // 最大生存周期，天
			// MaxBackups: 0, // 保留旧文件的最大个数
			LocalTime: true,  // 日志文件名时间格式为本地时间
			Compress:  false, // 是否压缩
		},
		"",
		log.LstdFlags,
	).WithCaller(2)

	return nil
}
