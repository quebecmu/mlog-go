package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/xxandjg/mlog-go/web/model"
	"go.uber.org/zap"
)

var Conf = new(model.App)

func Init(m string) error {
	viper.SetConfigFile("./conf/app-" + m + ".yaml")

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		zap.L().Sugar().Info("配置文件产生变化")
		if err2 := viper.Unmarshal(&Conf); err2 != nil {
			zap.L().Sugar().Errorf("unmarshal to Conf failed, err:%v", err2)
			panic(fmt.Errorf("unmarshal to Conf failed, err:%v", err2))
		}
	})

	err := viper.ReadInConfig()
	if err != nil {
		zap.L().Sugar().Errorf("ReadInConfig failed, err: %v", err)
		panic(fmt.Errorf("ReadInConfig failed, err: %v", err))
	}
	if err := viper.Unmarshal(&Conf); err != nil {
		zap.L().Sugar().Errorf("unmarshal to Conf failed, err:%v", err)
		panic(fmt.Errorf("unmarshal to Conf failed, err:%v", err))
	}
	return err

}
