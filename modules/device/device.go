package device

import (
	"github.com/gin-gonic/gin"
	"github.com/taerc/ezgo"
)

var M string = "DEVICE"

type QueryDevice struct {
	ezgo.GinFlow
}

func (qd *QueryDevice) Proc(ctx *gin.Context) {
	ezgo.Info(ctx, M, "query")
	qd.ResponseJson(ctx, 200, map[string]int{"id": 20, "name": 22})
}

func init() {
	route := ezgo.Group("/maicro/device")
	ezgo.SetPostProc(route, "query", &QueryDevice{})
}
