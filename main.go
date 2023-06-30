package main

import (
	"expezgo/modules/dblic"
	"expezgo/modules/device"
	"expezgo/modules/user"
	"github.com/taerc/ezgo"
	"github.com/taerc/ezgo/conf"
	_ "gorm.io/driver/mysql"
)

var M string = "MAIN"

func Init(data interface{}) int {

	conf.LoadConfigure(ezgo.ConfigPath)
	ezgo.LoadComponent(
		//ezgo.WithComponentLogger(conf.Config),
		ezgo.WithComponentMySQL(conf.Config),
	)

	ezgo.LoadModule(
		device.WithModuleDevice(),
		dblic.WithModuleLicence(),
		user.WithModuleUser(),
	)
	return ezgo.Success
}

func Exec(data interface{}) int {

	ezgo.Run("127.0.0.1:8080")
	return ezgo.Success

}

func Done(data interface{}) int {

	return ezgo.Success
}

func init() {
	ezgo.InitAppFlow(Init, Exec, Done)
}

func main() {

	ezgo.Do(nil)
}
