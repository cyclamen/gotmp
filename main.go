package main

import (
	"github.com/cyclamen/gotmp/conf"
	. "github.com/cyclamen/gotmp/log"
	"github.com/koding/multiconfig"
)

type flags struct {
	Conf     string
	LogLevel string
}

func main() {

	// check flag

	flagsLoader := &multiconfig.FlagLoader{}

	theFlag := &flags{}

	err := flagsLoader.Load(theFlag)

	if err != nil {
		panic(err)
	}

	err = conf.InitConfig(theFlag.Conf)
	if err != nil {
		panic(err)
	}

	var logLevel string
	if theFlag.LogLevel == "" {
		logLevel = conf.Conf.LogLevel
	} else {
		logLevel = theFlag.LogLevel
	}

	InitLog("NewAPP", 500, logLevel, true)

	Log.Info("Configure file init finished")
	Log.Info("Configure release mode is %v", conf.Conf.ReleaseMode)
	Log.Info("Configure log level is %s", conf.Conf.LogLevel)
	Log.Info("Configure app name is %s", conf.Conf.App.Name)
	Log.Info("Configure app backup number is %v", conf.Conf.App.NeedLog)
	Log.Info("Configure app backup number is %d", conf.Conf.App.BackupNumber)

}
