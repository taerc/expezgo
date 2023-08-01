package controller

import (
	"expezgo/modules/ent"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"github.com/taerc/ezgo"
	"sync"
)

type ServiceLPR struct {
}
type LPR struct {
	FiveYear decimal.Decimal `json:"five_year"`
	OneYear  decimal.Decimal `json:"one_year"`
	Date     string          `json:"date"`
}

func (s *ServiceLPR) Create(ctx *gin.Context) {

	lpr := LPR{}
	if e := ezgo.JsonBind(ctx, &lpr); e != nil {
		ezgo.ErrorResponse(ctx, e)
		return
	}

	if _, e := ent.DB().Debug().LPR.Create().
		SetCreateAt(ezgo.GetUnixTimeStamp()).
		SetFiveYear(lpr.FiveYear).
		SetOneYear(lpr.OneYear).
		SetUpdateAt(ezgo.GetUnixTimeStamp()).
		SetDate(lpr.Date).Save(ctx.Request.Context()); e != nil {
		ezgo.ErrorResponse(ctx, ezgo.NewEError(100004, e))
		return
	}

	ezgo.OKResponse(ctx, lpr)

}

func WithModuleLPR() func(wg *sync.WaitGroup) {

	return func(wg *sync.WaitGroup) {
		defer wg.Done()
		s := new(ServiceLPR)
		route := ezgo.Group("/maicro/lpr")
		ezgo.POST(route, "add", s.Create)
		ezgo.Info(nil, "Controller", "Load finished!")
	}

}
