package controller

import (
	"legend/model"
	"legend/service"
	"legend/tool"

	"github.com/gin-gonic/gin"
)

type ShengYuController struct {
}

func (syc *ShengYuController) Router(engine *gin.Engine) {
	engine.GET("/api/shengyu/query",syc.ShengYuQuery)
	engine.PUT("/api/shengyu/refresh/update",syc.ShengYuRefreshUpdate)
}

func (syc *ShengYuController) ShengYuQuery(ctx *gin.Context) {
	//调用service
	shengYuService := service.ShengYuService{}
	shengYus,err := shengYuService.ShengYuQueryService()
	if err != nil {
		tool.Failed(ctx,"获取数据失败")
		return
	}
	tool.Success(ctx,shengYus)
}

func (syc *ShengYuController) ShengYuRefreshUpdate(ctx *gin.Context) {
	var data model.ShengYu
	err := tool.Decode(ctx.Request.Body,&data)
	if err != nil {
		tool.Failed(ctx,"参数解析失败")
		return
	}
	//调用service
	shengYuService := service.ShengYuService{}
	res,err := shengYuService.ShengYuRefreshUpdateService(&data)
	if res <= 0 || err != nil {
		tool.Failed(ctx,"更新数据失败")
	}
	tool.Success(ctx,res)
}