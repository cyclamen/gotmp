package conf

import (
	"errors"
	"os"

	"github.com/koding/multiconfig"
	"log"
)

var (
	Conf              config
	defaultConfigFile = "conf/conf.toml"
)

type config struct {
	ReleaseMode bool   `toml:"release_mode"`
	LogLevel    string `toml:"log_level"`

	// 应用配置
	App app
}

type app struct {
	Name         string `toml:"name"`
	NeedLog      bool   `toml:"need_log"`
	BackupNumber int8   `toml:"backup_number"`
}

func InitConfig(configFile string) error {
	if configFile == "" {
		configFile = defaultConfigFile
	}

	Conf = config{
		ReleaseMode: false,
		LogLevel:    "DEBUG",
	}

	if _, err := os.Stat(configFile); err != nil {

		return errors.New("config file err:" + err.Error())
	} else {
		log.Print("Load configure from file " + configFile)
		configLoader := multiconfig.TOMLLoader{Path: configFile}

		err := configLoader.Load(&Conf)

		if err != nil {

			log.Fatal("Load configure file err:" + err.Error())

		}

	}

	return nil

}
