package main

import (
	"expezgo/pkg/controller"
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
		ezgo.WithComponentMySQL(ezgo.Default, &conf.Config.SQL),
	)

	ezgo.LoadModule(
		ezgo.WithModuleGitLab(),
		controller.WithModuleDevice(),
		controller.WithModuleLicence(),
		controller.WithModuleUser(),
		controller.WithModuleSwagger(),
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
