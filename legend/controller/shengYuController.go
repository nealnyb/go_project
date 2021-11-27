package controller

import (
	"legend/service"
	"legend/tool"

	"github.com/gin-gonic/gin"
)

type ShengYuController struct {
}

func (syc *ShengYuController) Router(engine *gin.Engine) {
	engine.GET("/api/shengyu/query",syc.ShengYuQuery)
}

func (sys *ShengYuController) ShengYuQuery(ctx *gin.Context) {
	//调用service
	shengYuService := service.ShengYuService{}
	shengYus,err := shengYuService.ShengYuQueryService()
	if err != nil {
		tool.Failed(ctx,"获取数据失败")
		return
	}
	tool.Success(ctx,shengYus)
}