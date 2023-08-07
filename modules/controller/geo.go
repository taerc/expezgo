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

type GEO struct {
}

func (g *GEO) List(ctx *gin.Context) {
	db, _ := ezgo.DB()
	pros := make([]Province, 0)
	//db.Find(&pros)
	//db.Preload("Cities").Find(&pros)
	db.Preload("Cities").Preload("Cities.Counties").Find(&pros)
	//db.Preload("Cities", "id & 1").Preload("Cities.Counties", "id &1 ").Find(&pros)
	ezgo.OKResponse(ctx, pros)

}

func WithModuleGEO() func(wg *sync.WaitGroup) {
	return func(wg *sync.WaitGroup) {
		defer wg.Done()
		s := new(GEO)
		route := ezgo.Group("/maicro/geo")
		ezgo.GET(route, "query", s.List)
		ezgo.Info(nil, "GEO", "Load finished!")
	}
}
