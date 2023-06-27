package main

import (
	_ "expezgo/modules/device"
	_ "expezgo/modules/user"
	log "github.com/sirupsen/logrus"
	"github.com/taerc/ezgo"
)

func Init(data interface{}) int {

	log.Info("Init")
	//device.InitDevRouter()
	return ezgo.Success
}

func Exec(data interface{}) int {

	log.Info("Execute")
	ezgo.Run("127.0.0.1:8080")
	return ezgo.Success

}

func Done(data interface{}) int {

	log.Info("Done")

	return ezgo.Success
}

var appFlow *ezgo.AppFlow = nil

func init() {
	appFlow = ezgo.InitAppFlow(Init, Exec, Done)
}

func main() {
	log.Info("hello")
	appFlow.Do(nil)
}
