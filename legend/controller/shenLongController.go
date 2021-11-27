package controller

import (
	"legend/model"
	"legend/service"
	"legend/tool"

	"github.com/gin-gonic/gin"
)

type ShenLongController struct {
}

func (slc *ShenLongController) Router(engine *gin.Engine){
	engine.GET("/api/shenlong/query",slc.ShenLongQuery)
	engine.PUT("/api/shenlong/refresh/update",slc.ShenLongRefreshUpdate)
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

func (slc *ShenLongController) ShenLongRefreshUpdate(ctx *gin.Context) {
	var data model.ShenLong
	err := tool.Decode(ctx.Request.Body,&data)
	if err != nil {
		tool.Failed(ctx,"参数解析失败")
		return
	}
	//调用service
	shenLongService := service.ShenLongService{}
	res,err := shenLongService.ShenLongRefreshUpdateService(&data)
	if res <= 0 || err != nil {
		tool.Failed(ctx,"更新数据失败")
	}
	tool.Success(ctx,res)
}