package config

import (
	"os"
	"sync"
	"github.com/spf13/viper"
	"github.com/sirupsen/logrus"
)

type ConfigManager struct {
	SourceType				string
	ChannelAction		string
	SinkAction			string
}

type ConfigManagerBehaviour interface{
	LogConfig()
}

func (config *ConfigManager) LogConfig() {
	logrus.WithFields(logrus.Fields{
		"SourceType":									ConfigManagerInstance.SourceType,
		"ChannelAction":							ConfigManagerInstance.ChannelAction,
		"SinkAction":									ConfigManagerInstance.SinkAction,
	}).Info("configurationManager loaded")
}

var once sync.Once
var ConfigManagerInstance *ConfigManager

func setupConfigPath() {
	viper.SetConfigName("config")
	configPath, exist := os.LookupEnv("CONFIG_PATH")
	if exist {
		viper.AddConfigPath(configPath)
	}
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

func NewConfigManager() ConfigManagerBehaviour {
	once.Do(func() {
		setupConfigPath()

		sourceType := viper.GetString("source.type")
		channelAction := viper.GetString("channel.action")
		sinkAction := viper.GetString("sink.action")

		ConfigManagerInstance = &ConfigManager{
			SourceType:			sourceType,
			ChannelAction:	channelAction,
			SinkAction:			sinkAction,
		}

		ConfigManagerInstance.LogConfig()
	})

	return ConfigManagerInstance
}
