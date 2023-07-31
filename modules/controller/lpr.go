package controller

import (
	"expezgo/modules/ent"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"github.com/taerc/ezgo"
	"sync"
)

type LPR struct {
	*ezgo.GinFlow
	FiveYear decimal.Decimal `json:"five_year"`
	OneYear  decimal.Decimal `json:"one_year"`
	Date     string          `json:"date"`
}

func (l *LPR) Proc(ctx *gin.Context) {

	if e := l.Bind(ctx, l); e != nil {
		l.Error(ctx, e)
		return
	}

	if _, e := ent.DB().Debug().LPR.Create().
		SetCreateAt(ezgo.GetUnixTimeStamp()).
		SetFiveYear(l.FiveYear).
		SetOneYear(l.OneYear).
		SetUpdateAt(ezgo.GetUnixTimeStamp()).
		SetDate(l.Date).Save(ctx.Request.Context()); e != nil {
		l.Error(ctx, ezgo.NewEError(100004, e))
		return
	}

	l.Response(ctx, l)
}

func WithModuleLPR() func(wg *sync.WaitGroup) {

	return func(wg *sync.WaitGroup) {
		route := ezgo.Group("/maicro/lpr")
		ezgo.ProcPOST(route, "add", &LPR{})
		ezgo.Info(nil, M, "Load finished!")
		wg.Done()
	}

}
