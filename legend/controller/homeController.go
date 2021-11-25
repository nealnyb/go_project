package controller

import (
	"legend/service"
	"legend/tool"
	"github.com/gin-gonic/gin"
)

type HomeController struct {
}

func (hc *HomeController) Router(engine *gin.Engine) {
	engine.GET("/api/home/query",hc.HomeQuery)	
}

func (hc *HomeController) HomeQuery(ctx *gin.Context){
	//调用service
	homeService := &service.HomeService{}
	homes,err := homeService.HomeQueryService()
	if err != nil {
		tool.Failed(ctx,"获取数据失败")
		return
	}
	tool.Success(ctx,homes)
}