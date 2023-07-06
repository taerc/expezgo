package main

import (
	"expezgo/modules/dblic"
	"expezgo/modules/device"
	"expezgo/modules/doc"
	"expezgo/modules/user"
	"fmt"
	"github.com/taerc/ezgo"
	"github.com/taerc/ezgo/conf"
	_ "gorm.io/driver/mysql"
)

var M string = "MAIN"

func Init(data interface{}) int {

	conf.LoadConfigure(ezgo.ConfigPath)
	ezgo.LoadComponent(
		ezgo.WithComponentResource(conf.Config),
		//ezgo.WithComponentLogger(conf.Config),
		ezgo.WithComponentMySQL(ezgo.Default, conf.Config),
	)

	ezgo.LoadModule(
		ezgo.WithModuleGitLab(),
		device.WithModuleDevice(),
		dblic.WithModuleLicence(),
		user.WithModuleUser(),
		doc.WithModuleSwagger(),
	)
	return ezgo.Success
}

func Exec(data interface{}) int {

	ezgo.Run(fmt.Sprintf("%s:%s", conf.Config.Host, conf.Config.Port))
	return ezgo.Success

}

func Done(data interface{}) int {

	return ezgo.Success
}

func init() {
	ezgo.InitAppFlow(Init, Exec, Done)
}

//go:generate swag init
func main() {

	ezgo.Do(nil)
}
