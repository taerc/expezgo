package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/taerc/ezgo"
	"sync"
)

type Province struct {
	ProvinceId uint32 `gorm:"primaryKey; column:id" json:"province_id"`
	Name       string `gorm:"column:name" json:"name"`
	Type       int    `gorm:"column:name" json:"type"`
	Cities     []City `json:"cities" gorm:"foreignKey:Pid;references:ProvinceId"`
}

func (p Province) TableName() string {
	return "province"
}

type City struct {
	CityId   uint32   `gorm:"primaryKey; column:id" json:"city_id"`
	Name     string   `gorm:"column:name" json:"name"`
	Type     int      `gorm:"column:type" json:"type"`
	Pid      uint32   `gorm:"column:pid" json:"pid"`
	Counties []County `json:"counties" gorm:"foreignKey:Pid;references:CityId"`
}

func (c City) TableName() string {
	return "city"
}

type County struct {
	CountyId uint32 `gorm:"primaryKey; column:id" json:"county_id"`
	Name     string `gorm:"column:name" json:"name"`
	Type     int    `gorm:"column:type" json:"type"`
	Pid      uint32 `gorm:"column:pid" json:"pid"`
}

func (c County) TableName() string {
	return "county"
}

type QueryGEO struct {
	*ezgo.GinFlow
}

func queryGEO() {

}

func (geo *QueryGEO) Proc(ctx *gin.Context) {
	//queryGEO()
	db := ezgo.DB()
	pros := make([]Province, 0)
	//db.Find(&pros)
	//db.Preload("Cities").Find(&pros)
	//db.Preload("Cities").Preload("Cities.Counties").Find(&pros)
	db.Preload("Cities", "id & 1").Preload("Cities.Counties", "id &1 ").Find(&pros)

	geo.ResponseIndJson(ctx, ezgo.Success, pros)

	//if data, e := json.Marshal(pros); e == nil {
	//	fmt.Println(string(data))
	//}
	//fmt.Println(pros)
}

func WithModuleGEO() func(wg *sync.WaitGroup) {
	return func(wg *sync.WaitGroup) {
		wg.Done()
		route := ezgo.Group("/maicro/geo")
		ezgo.ProcGET(route, "query", &QueryGEO{})
		ezgo.Info(nil, M, "Load finished!")
	}
}
