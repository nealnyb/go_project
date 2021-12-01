package controller

import (
	"legend/model"
	"legend/service"
	"legend/tool"

	"github.com/gin-gonic/gin"
)

type ZhenYingController struct {
}

func (zyc *ZhenYingController) Rotuer(engine *gin.Engine){
	engine.GET("/zhenying/query",zyc.ZhenYingQuery)
	engine.POST("/zhenying/refresh/update",zyc.ZhenYingRefreshUpdate)
}

func (zyc *ZhenYingController) ZhenYingQuery(ctx *gin.Context) {
	//调用service
	zhenYingService := service.ZhenYingService{}
	zhenYings,err := zhenYingService.ZhenYingQueryService()
	if err != nil {
		tool.Failed(ctx,"获取数据失败")
	}
	tool.Success(ctx,zhenYings)
}

func (zyc *ZhenYingController) ZhenYingRefreshUpdate(ctx *gin.Context) {
	var data model.ZhenYing
	err := tool.Decode(ctx.Request.Body,&data)
	if err != nil {
		tool.Failed(ctx,"参数解析失败")
		return
	}
	//调用service
	zhenYingService := service.ZhenYingService{}
	res,err := zhenYingService.ZhenYingRefreshUpdateService(&data)
	if res <= 0 || err != nil {
		tool.Failed(ctx,"更新数据失败")
	}
	tool.Success(ctx,res)
}