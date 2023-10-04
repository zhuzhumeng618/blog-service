package setting

import (
	"time"
)

// Server 服务配置模型，与 config.yaml 配置项对应
type Server struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// App 应用配置模型，与 config.yaml 配置项对应
type App struct {
	DefaultPageSize int
	MaxPageSize     int
	LogSavePath     string
	LogFileName     string
	LogFileExt      string
}

// Database 数据源配置模型，与 config.yaml 配置项对应
type Database struct {
	DBType       string
	Username     string
	Password     string
	Host         string
	DBName       string
	TablePrefix  string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

// ReadSection 读取各节点配置
// Params: key 表示配置文件根节点名称，value 表示将配置文件序列化到哪个结构体实例中
func (s *Setting) ReadSection(key string, value interface{}) error {
	if err := s.vip.UnmarshalKey(key, value); err != nil {
		return err
	}
	return nil
}
