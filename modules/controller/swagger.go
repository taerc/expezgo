package controller

import (
	"expezgo/docs"
	"expezgo/version"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/taerc/ezgo"
	"sync"
)

// WithModuleSwagger 注册swagger路由 godoc
// @description swagger 必须子项目中自动生成
func WithModuleSwagger() func(wg *sync.WaitGroup) {

	return func(wg *sync.WaitGroup) {
		wg.Done()
		r := ezgo.Group("/docs")
		ezgo.GET(r, "/swagger/*any", func(ctx *gin.Context) {
			docs.SwaggerInfo.Host = ctx.Request.Host
			docs.SwaggerInfo.Title = "EasyGo"
			docs.SwaggerInfo.Version = version.AppVersion
			docs.SwaggerInfo.InfoInstanceName = "EasyGo"
			ginSwagger.WrapHandler(swaggerFiles.Handler)(ctx)
		})
	}

}
