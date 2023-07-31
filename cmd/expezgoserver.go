package main

import (
	controller2 "expezgo/modules/controller"
	"flag"
	"fmt"
	"github.com/taerc/ezgo"
	"github.com/taerc/ezgo/conf"
	_ "gorm.io/driver/mysql"
)

var M string = "MAIN"

var ConfigPath string
var ShowVersion bool

func init() {
	flag.BoolVar(&ShowVersion, "version", false, "print program build version")
	flag.StringVar(&ConfigPath, "c", "conf/config.toml", "path of configure file.")
	flag.Parse()
}

func Init(data interface{}) int {

	conf.LoadConfigure(ConfigPath)
	ezgo.LoadComponent(
		ezgo.WithComponentResource(conf.Config),
		//ezgo.WithComponentLogger(conf.Config),
		ezgo.WithComponentMySQL(ezgo.Default, &conf.Config.SQL),
	)

	ezgo.LoadModule(
		ezgo.WithModuleGitLab(),
		controller2.WithModuleDevice(),
		controller2.WithModuleLicence(),
		controller2.WithModuleUser(),
		controller2.WithModuleGEO(),
		controller2.WithModuleSwagger(),
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
