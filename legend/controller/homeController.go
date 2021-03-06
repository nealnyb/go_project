package controller

import (
	"legend/model"
	"legend/service"
	"legend/tool"
	"github.com/gin-gonic/gin"
)

type HomeController struct {
}

func (hc *HomeController) Router(engine *gin.Engine) {
	engine.GET("/home/query",hc.HomeQuery)
	engine.POST("/home/refresh/update",hc.HomeRefeshUpdate)	
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

func (hc *HomeController) HomeRefeshUpdate(ctx *gin.Context){
	var data model.Home
	err := tool.Decode(ctx.Request.Body,&data)
	if err != nil {
		tool.Failed(ctx,"参数解析失败")
		return
	}
	//调用service
	homeService := &service.HomeService{}
	res,err := homeService.HomeRefeshUpdateService(&data)
	if res <= 0 || err != nil {
		tool.Failed(ctx,"数据更新失败")
		return
	}
	tool.Success(ctx,res)
}