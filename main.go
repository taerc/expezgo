package main

import (
	_ "expezgo/modules/device"
	"github.com/taerc/ezgo"
	"github.com/taerc/ezgo/conf"
	_ "gorm.io/driver/mysql"
)

var M string = "main"

func Init(data interface{}) int {

	ezgo.Info(nil, M, "Init")
	ezgo.Info(nil, M, "Hello")
	ezgo.Info(nil, M, ezgo.ConfigPath)
	conf.LoadConfigure(ezgo.ConfigPath)
	ezgo.Info(nil, M, conf.Config.LogDir)
	ezgo.Info(nil, M, conf.Config.Host)
	ezgo.Info(nil, M, conf.Config.LogFileName)
	ezgo.LoadComponent(ezgo.WithComponentLogger(conf.Config),
		ezgo.WithComponentMySQL(conf.Config),
	)
	return ezgo.Success
}

func Exec(data interface{}) int {

	ezgo.Info(nil, M, "Exec")
	ezgo.Run("127.0.0.1:8080")
	return ezgo.Success

}

func Done(data interface{}) int {

	ezgo.Info(nil, M, "Done")

	return ezgo.Success
}

var appFlow *ezgo.AppFlow = nil

func init() {
	appFlow = ezgo.InitAppFlow(Init, Exec, Done)
}

func main() {

	appFlow.Do(nil)
}
