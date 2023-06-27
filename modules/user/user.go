package user

import (
	"github.com/gin-gonic/gin"
	"github.com/taerc/ezgo"
)

type QueryUser struct {
	ezgo.GinFlow
}

func (qd *QueryUser) Proc(ctx *gin.Context) {
	qd.ResponseJson(ctx, 200, map[string]int{"age": 66, "name": 22})
}

func init() {
	route := ezgo.Group("/maicro/user")
	ezgo.SetPostProc(route, "query", &QueryUser{})
}
