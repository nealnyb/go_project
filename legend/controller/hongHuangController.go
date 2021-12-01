package controller

import (
	"legend/model"
	"legend/service"
	"legend/tool"

	"github.com/gin-gonic/gin"
)

type HongHuangController struct {
}

func (hhc *HongHuangController) Router(engine *gin.Engine) {
	engine.GET("/honghuang/query",hhc.HongHuangQuery)
	engine.POST("/honghuang/refresh/update",hhc.HongHuangRefreshUpdate)
}

func (hhc *HongHuangController) HongHuangQuery(ctx *gin.Context){
	//调用service
	hongHuangService := &service.HongHuangService{}
	hongHuangs,err := hongHuangService.HongHuangQueryService()
	if err != nil {
		tool.Failed(ctx,"获取数据失败")
		return
	}
	tool.Success(ctx,hongHuangs)
}

func (hcc *HongHuangController) HongHuangRefreshUpdate(ctx *gin.Context){
	var data model.HongHuang
	err := tool.Decode(ctx.Request.Body,&data)
	if err != nil {
		tool.Failed(ctx,"参数解析失败")
		return
	}
	//调用service
	hongHuangService := &service.HongHuangService{}
	res,err := hongHuangService.HongHuangRefreshUpdateService(&data)
	if res <= 0 || err != nil {
		tool.Failed(ctx,"更新数据失败")
		return
	}
	tool.Success(ctx,res)
}