package user

import (
	"github.com/gin-gonic/gin"
	"github.com/taerc/ezgo"
	"sync"
)

var M string = "USER"

type QueryUser struct {
	ezgo.GinFlow
}

func UserGroup(ctx *gin.Context) {

	ctx.JSON(ezgo.Success, ezgo.Response{
		Code:    10,
		Data:    map[string]int{"group": 66},
		Message: "ok",
	})
}

// Proc 用户信息查询 godoc
// @Summary 用户信息查询
// @Schemes
// @Description
// @Tags 用户管理
// @Accept json
// @Produce json
// @Success 200 {string} OK
// @Router /maicro/user/query [Post]
func (qd *QueryUser) Proc(ctx *gin.Context) {
	ezgo.Info(ctx, M, "query")
	qd.ResponseJson(ctx, 200, map[string]int{"age": 66, "name": 22})
}

type CheckUser struct {
	ezgo.GinFlow
}

// Proc 用户检查 godoc
// @Summary 用户信息检查
// @Schemes
// @Description
// @Tags 用户管理
// @Accept json
// @Produce json
// @Success 200 {string} OK
// @Router /maicro/user/check [Post]
func (qu *CheckUser) Proc(ctx *gin.Context) {
	ezgo.Info(ctx, M, "check")
	qu.ResponseJson(ctx, 200, map[string]int{"check": 99, "id": 8})
}

func init() {
}

func WithModuleUser() func(wg *sync.WaitGroup) {

	return func(wg *sync.WaitGroup) {
		wg.Done()
		route := ezgo.Group("/maicro/user", UserGroup)
		ezgo.SetPostProc(route, "query", &QueryUser{})
		ezgo.SetPostProc(route, "check", &CheckUser{})
		ezgo.Info(nil, M, "Load finished!")
	}
}
