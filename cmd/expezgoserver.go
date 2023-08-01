package main

import (
	"expezgo/modules/controller"
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

func Init(data interface{}) error {

	conf.LoadConfigure(ConfigPath)
	ezgo.LoadComponent(
		ezgo.WithComponentResource(conf.Config),
		//ezgo.WithComponentLogger(conf.Config),
		ezgo.WithComponentMySQL(ezgo.Default, &conf.Config.SQL),
	)

	ezgo.LoadModule(
		ezgo.WithModuleGitLab(),
		controller.WithModuleLPR(),
		controller.WithModuleSwagger(),
	)
	return nil
}

func Exec(data interface{}) error {

	ezgo.Run(fmt.Sprintf("%s:%s", conf.Config.Host, conf.Config.Port))
	return nil

}

func Done(data interface{}) error {

	return nil
}

func init() {
	ezgo.InitApplication(Init, Exec, Done)
}

//go:generate swag init
func main() {

	ezgo.Do(nil)
}
