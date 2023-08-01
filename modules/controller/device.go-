package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/taerc/ezgo"
	"sync"
)

type QueryDevice struct {
	ezgo.GinFlow
}

// Proc query device godoc
// @Summary query device
// @Schemes
// @Description
// @Tags 设备管理
// @Accept json
// @Produce json
// @Success 200 {string} OK
// @Router /maicro/device/query [Post]
func (qd *QueryDevice) Proc(ctx *gin.Context) {
	ezgo.Info(ctx, M, "query")
	qd.ResponseJson(ctx, 200, map[string]int{"id": 20, "name": 22})
}

func init() {
}

func WithModuleDevice() func(wg *sync.WaitGroup) {

	return func(wg *sync.WaitGroup) {
		route := ezgo.Group("/maicro/device")
		ezgo.ProcPOST(route, "query", &QueryDevice{})
		ezgo.Info(nil, M, "Load finished!")
		wg.Done()
	}

}
