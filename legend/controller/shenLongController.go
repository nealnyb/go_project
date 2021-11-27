package controller

import (
	"legend/service"
	"legend/tool"

	"github.com/gin-gonic/gin"
)

type ShenLongController struct {
}

func (slc *ShenLongController) Router(engine *gin.Engine){
	engine.GET("/api/shenlong/query",slc.ShenLongQuery)
}

func (slc *ShenLongController) ShenLongQuery(ctx *gin.Context) {
	//调用service
	shenLongService := service.ShenLongService{}
	shenLongs,err := shenLongService.ShenLongQueryService()
	if err != nil {
		tool.Failed(ctx,"获取数据失败")
	}
	tool.Success(ctx,shenLongs)
}