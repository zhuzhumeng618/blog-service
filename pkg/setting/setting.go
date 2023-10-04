package setting

import (
	"github.com/spf13/viper"
)

// Setting 配置模型，封装 viper
type Setting struct {
	vip *viper.Viper
}

// NewSetting 新建配置实例
func NewSetting() (*Setting, error) {
	vip := viper.New()
	// 设置配置文件名称
	vip.SetConfigName("config")
	// 添加配置文件路径
	vip.AddConfigPath("configs/")
	// 设置配置文件类型
	vip.SetConfigType("yaml")
	// 读取配置文件
	if err := vip.ReadInConfig(); err != nil {
		return nil, err
	}
	// 返回配置文件实例
	return &Setting{vip: vip}, nil
}
