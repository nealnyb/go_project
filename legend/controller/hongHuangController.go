package controller

import (
	"legend/service"
	"legend/tool"
	"github.com/gin-gonic/gin"
)

type HongHuangController struct {
}

func (hhc *HongHuangController) Router(engine *gin.Engine) {
	engine.GET("/api/honghuang/query",hhc.HongHuangQuery)
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