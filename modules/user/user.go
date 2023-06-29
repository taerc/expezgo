package user

import (
	"github.com/gin-gonic/gin"
	"github.com/taerc/ezgo"
)

var M string = "USER"

type QueryUser struct {
	ezgo.GinFlow
}

func UserGroup(ctx *gin.Context) {

	ezgo.Info(ctx, M)

	ctx.JSON(ezgo.Success, ezgo.ResponseTemplate{
		Code:    10,
		Data:    map[string]int{"group": 66},
		Message: "ok",
	})
}

func (qd *QueryUser) Proc(ctx *gin.Context) {
	ezgo.Info(ctx, M, "query")
	qd.ResponseJson(ctx, 200, map[string]int{"age": 66, "name": 22})
}

type CheckUser struct {
	ezgo.GinFlow
}

func (qu *CheckUser) Proc(ctx *gin.Context) {
	ezgo.Info(ctx, M, "check")
	qu.ResponseJson(ctx, 200, map[string]int{"check": 99, "id": 8})
}

func init() {
	route := ezgo.Group("/maicro/user", UserGroup)
	ezgo.SetPostProc(route, "query", &QueryUser{})
	ezgo.SetPostProc(route, "check", &CheckUser{})
}
