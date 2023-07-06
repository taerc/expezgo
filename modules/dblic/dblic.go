package dblic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/taerc/ezgo"
	"sync"
)

var M string = "DBLIC"

type LicenceSN struct {
	Id         int64  `gorm:"primary_key; column:id" json:"id"`
	DevUUID    string `gorm:"column:dev_uuid" json:"dev_uuid"`
	LicPath    string `gorm:"column:lic_path" json:"lic_path"`
	State      int    `gorm:"column:state" json:"state"`
	TaskId     int64  `gorm:"column:task_id" json:"task_id"`
	CreateTime uint64 `gorm:"column:create_time" json:"create_time"`
}

func queryLicence() []LicenceSN {
	db := ezgo.DB()

	lic := make([]LicenceSN, 0)
	db.Table("lic_sn").Where("id > 0").Scan(&lic)

	for _, l := range lic {
		fmt.Println(l.DevUUID)
	}
	return lic
}

type QueryLic struct {
	*ezgo.GinFlow
}

func (ql *QueryLic) Proc(ctx *gin.Context) {

	ls := queryLicence()
	ql.ResponseJson(ctx, 200, ls)
}

func WithModuleLicence() func(wg *sync.WaitGroup) {

	return func(wg *sync.WaitGroup) {
		wg.Done()
		route := ezgo.Group("/maicro/lic")
		ezgo.SetPostProc(route, "query", &QueryLic{})
		ezgo.Info(nil, M, "Load finished!")
	}

}
