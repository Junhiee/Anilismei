package configs

import (
	"git.virjar.com/Junhiee/anilismei/pkg/log"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type DateBase struct{}

type Server struct{}

type Logger struct{}

type Config struct {
	*DateBase
	*Server
}

var cfg Config

func InitConfig() {
	viper.SetConfigName("conf")      // 文件名
	viper.SetConfigType("toml")      // 文件类型
	viper.AddConfigPath("./configs") // 配置文件路径
	err := viper.ReadInConfig()
	if err != nil {
		log.ZLOG.Panic("Viper config init err", zap.Error(err))
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		log.ZLOG.Panic("Viper unmarshal faild", zap.Error(err))
	}
}
