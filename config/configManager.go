package config

import (
	"os"
	"sync"
	"github.com/spf13/viper"
	"github.com/sirupsen/logrus"
)

type ConfigManager struct {
	IngestorType		string
	SourcePath			string
	ProcessorAction	string
	FormatterType		string
}

type ConfigManagerBehaviour interface{
	LogConfig()
}

func (config *ConfigManager) LogConfig() {
	logrus.WithFields(logrus.Fields{
		"IngestorType":									ConfigManagerInstance.IngestorType,
		"ProcessorAction":							ConfigManagerInstance.ProcessorAction,
		"FormatterType":								ConfigManagerInstance.FormatterType,
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

		ingestorType := viper.GetString("ingestor.type")
		sourcePath := viper.GetString("ingestor.sourcePath")
		channelAction := viper.GetString("processor.action")
		formatterType := viper.GetString("formatter.type")

		ConfigManagerInstance = &ConfigManager{
			IngestorType:			ingestorType,
			SourcePath:				sourcePath,
			ProcessorAction:	channelAction,
			FormatterType:		formatterType,
		}

		ConfigManagerInstance.LogConfig()
	})

	return ConfigManagerInstance
}
